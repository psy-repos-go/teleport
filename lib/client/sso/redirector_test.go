/*
 * Teleport
 * Copyright (C) 2024  Gravitational, Inc.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Affero General Public License for more details.
 *
 * You should have received a copy of the GNU Affero General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package sso_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/teleport/api/utils/keys"
	"github.com/gravitational/teleport/lib/client/sso"
	"github.com/gravitational/teleport/lib/web"
)

func TestRedirector(t *testing.T) {
	ctx := context.Background()
	username := "alice"

	mockProxy := newMockProxy(t)

	// Create a basic redirector.
	rd, err := sso.NewRedirector(sso.RedirectorConfig{
		ProxyAddr: mockProxy.URL,
	})
	require.NoError(t, err)
	t.Cleanup(rd.Close)

	// Ensure that ClientCallbackURL is a valid url.
	_, err = url.Parse(rd.ClientCallbackURL)
	require.NoError(t, err)

	// Construct a fake ssh login response with the redirector's client callback URL.
	successResponseURL, err := web.ConstructSSHResponse(web.AuthParams{
		ClientRedirectURL: rd.ClientCallbackURL,
		Username:          username,
	})
	require.NoError(t, err)

	newErrorResponseURL := func(err error) string {
		failureResponseURL, _ := url.Parse(rd.ClientCallbackURL)
		query := failureResponseURL.Query()
		query.Set("err", err.Error())
		failureResponseURL.RawQuery = query.Encode()
		return failureResponseURL.String()
	}

	for _, tt := range []struct {
		name             string
		idpHandler       http.HandlerFunc
		privateKeyPolicy keys.PrivateKeyPolicy
		expectRedirect   string
		assertErr        require.ErrorAssertionFunc
	}{
		{
			name: "OK login success",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, successResponseURL.String(), http.StatusPermanentRedirect)
			},
			expectRedirect: sso.LoginSuccessRedirectURL,
			assertErr:      require.NoError,
		},
		{
			name: "NOK no login response",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				// No response or error encoded.
				http.Redirect(w, r, rd.ClientCallbackURL, http.StatusPermanentRedirect)
			},
			expectRedirect: sso.LoginFailedRedirectURL,
			assertErr:      require.Error,
		},
		{
			name: "NOK server error",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				// Encode a login error in the client callback URL.
				http.Redirect(w, r, newErrorResponseURL(errors.New("login failed")), http.StatusPermanentRedirect)
			},
			expectRedirect: sso.LoginFailedRedirectURL,
			assertErr: func(t require.TestingT, err error, v ...any) {
				require.ErrorContains(t, err, "login failed", "expected login failed error but got %v", err)
			},
		},
		{
			name: "NOK indirect server failure",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				// The client may redirect straight to the proxy if the client callback is misformed
				// or from other indirect login failures.
				proxyRedirectURL, err := url.Parse(mockProxy.URL)
				require.NoError(t, err)

				proxyRedirectURL.Path = sso.LoginFailedBadCallbackRedirectURL
				http.Redirect(w, r, proxyRedirectURL.String(), http.StatusPermanentRedirect)
			},
			expectRedirect: sso.LoginFailedBadCallbackRedirectURL,
			assertErr: func(t require.TestingT, err error, v ...any) {
				// The sso login will timeout due to the client callback never being redirected to.
				require.ErrorIs(t, err, context.DeadlineExceeded)
			},
		},
		// PrivateKeyPolicy tests
		{
			name: "OK close redirect failed hardware_key login",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				// Encode a private key policy error in the client callback URL.
				err := keys.NewPrivateKeyPolicyError(keys.PrivateKeyPolicyHardwareKey)
				http.Redirect(w, r, newErrorResponseURL(err), http.StatusPermanentRedirect)
			},
			expectRedirect: sso.LoginClose,
			assertErr: func(tt require.TestingT, err error, i ...any) {
				policy, err := keys.ParsePrivateKeyPolicyError(err)
				require.NoError(t, err, "expected private key policy error but got %v", err)
				require.Equal(t, keys.PrivateKeyPolicyHardwareKey, policy)
			},
		},
		{
			name: "OK close redirect failed hardware_key_touch login",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				// Encode a private key policy error in the client callback URL.
				err := keys.NewPrivateKeyPolicyError(keys.PrivateKeyPolicyHardwareKeyTouch)
				http.Redirect(w, r, newErrorResponseURL(err), http.StatusPermanentRedirect)
			},
			expectRedirect: sso.LoginClose,
			assertErr: func(tt require.TestingT, err error, i ...any) {
				policy, err := keys.ParsePrivateKeyPolicyError(err)
				require.NoError(t, err, "expected private key policy error but got %v", err)
				require.Equal(t, keys.PrivateKeyPolicyHardwareKeyTouch, policy)
			},
		},
		{
			name: "OK terminal redirect on failed hardware_key_pin login",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				// Encode a private key policy error in the client callback URL.
				err := keys.NewPrivateKeyPolicyError(keys.PrivateKeyPolicyHardwareKeyPIN)
				http.Redirect(w, r, newErrorResponseURL(err), http.StatusPermanentRedirect)
			},
			expectRedirect: sso.LoginTerminalRedirectURL,
			assertErr: func(tt require.TestingT, err error, i ...any) {
				policy, err := keys.ParsePrivateKeyPolicyError(err)
				require.NoError(t, err, "expected private key policy error but got %v", err)
				require.Equal(t, keys.PrivateKeyPolicyHardwareKeyPIN, policy)
			},
		},
		{
			name: "OK terminal redirect on failed hardware_key_touch_and_pin login",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				// Encode a private key policy error in the client callback URL.
				err := keys.NewPrivateKeyPolicyError(keys.PrivateKeyPolicyHardwareKeyTouchAndPIN)
				http.Redirect(w, r, newErrorResponseURL(err), http.StatusPermanentRedirect)
			},
			expectRedirect: sso.LoginTerminalRedirectURL,
			assertErr: func(tt require.TestingT, err error, i ...any) {
				policy, err := keys.ParsePrivateKeyPolicyError(err)
				require.NoError(t, err, "expected private key policy error but got %v", err)
				require.Equal(t, keys.PrivateKeyPolicyHardwareKeyTouchAndPIN, policy)
			},
		},
		{
			name: "OK success redirect on success with hardware_key",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, successResponseURL.String(), http.StatusPermanentRedirect)
			},
			privateKeyPolicy: keys.PrivateKeyPolicyHardwareKey,
			expectRedirect:   sso.LoginSuccessRedirectURL,
			assertErr:        require.NoError,
		},
		{
			name: "OK success redirect on success with hardware_key_pin",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, successResponseURL.String(), http.StatusPermanentRedirect)
			},
			privateKeyPolicy: keys.PrivateKeyPolicyHardwareKeyPIN,
			expectRedirect:   sso.LoginSuccessRedirectURL,
			assertErr:        require.NoError,
		},
		{
			name: "OK terminal redirect on success with hardware_key_touch",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, successResponseURL.String(), http.StatusPermanentRedirect)
			},
			privateKeyPolicy: keys.PrivateKeyPolicyHardwareKeyTouch,
			expectRedirect:   sso.LoginTerminalRedirectURL,
			assertErr:        require.NoError,
		},
		{
			name: "OK terminal redirect on success with hardware_key_touch_and_pin",
			idpHandler: func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, successResponseURL.String(), http.StatusPermanentRedirect)
			},
			privateKeyPolicy: keys.PrivateKeyPolicyHardwareKeyTouchAndPIN,
			expectRedirect:   sso.LoginTerminalRedirectURL,
			assertErr:        require.NoError,
		},
	} {
		t.Run(tt.name, func(t *testing.T) {
			rd.PrivateKeyPolicy = tt.privateKeyPolicy

			// Open a mock IdP server which will handle a redirect and result in the expected IdP session payload.
			mockIdPServer := httptest.NewServer(tt.idpHandler)
			t.Cleanup(mockIdPServer.Close)

			// connecting to the mockIdPServer should redirect to the client callback, parsing the login response.
			// We should be redirected to the sso success page.
			resp, err := http.Get(mockIdPServer.URL)
			require.NoError(t, err)
			defer resp.Body.Close()

			body, err := io.ReadAll(resp.Body)
			require.NoError(t, err)
			require.Equal(t, tt.expectRedirect, string(body))

			// Sending a request to the IdP server should result in a redirector callback result.
			ctx, cancel := context.WithTimeout(ctx, time.Second)
			defer cancel()

			loginResponse, err := rd.WaitForResponse(ctx)
			tt.assertErr(t, err)

			if err == nil {
				require.Equal(t, username, loginResponse.Username)
			}
		})
	}
}

func TestValidateClientRedirect(t *testing.T) {
	t.Run("StandardLocalhost", func(t *testing.T) {
		for _, goodURL := range []string{
			"http://127.0.0.1/callback",
			"http://127.0.0.1:1234/callback",
			"http://[::1]/callback",
			"http://[::1]:1234/callback",
			"http://localhost/callback",
			"http://localhost:1234/callback",
		} {
			for name, ceremonyType := range map[string]sso.CeremonyType{
				"Login": sso.CeremonyTypeLogin,
				"MFA":   sso.CeremonyTypeMFA,
				"Test":  sso.CeremonyTypeTest,
			} {
				t.Run(fmt.Sprintf("Ceremony%v", name), func(t *testing.T) {
					var defaultSettings *types.SSOClientRedirectSettings
					require.NoError(t, sso.ValidateClientRedirect(goodURL+"?secret_key=", ceremonyType, defaultSettings))
				})
			}
		}
	})

	t.Run("InvalidLocalhostLike", func(t *testing.T) {
		for _, badURL := range []string{
			"http://127.0.0.1:12345/notcallback",
			"http://127a0.0.1/callback",
			"http://127a0.0.1:1234/callback",
			"http://::1/callback",
			"http://notlocalhost/callback",
			"http://sub.localhost/callback",
			"http://localhost.com/callback",
			"http://127.0.0.1.example.com/callback",
			"http://[::1].example.com/callback",
			"http://username@127.0.0.1:12345/callback",
			"http://@localhost:12345/callback",
			"http://localhost@example.com/callback",
			"http://127.0.0.1:12345@example.com/callback",
			"https://127.0.0.1:12345/callback",
			"https://localhost:12345/callback",
			"https://localhost/callback",
			"ftp://localhost/callback",
			"ftp://127.0.0.1/callback",
			"ftp://[::1]/callback",
		} {
			for name, ceremonyType := range map[string]sso.CeremonyType{
				"Login": sso.CeremonyTypeLogin,
				"MFA":   sso.CeremonyTypeMFA,
				"Test":  sso.CeremonyTypeTest,
			} {
				t.Run(fmt.Sprintf("Ceremony%v", name), func(t *testing.T) {
					var defaultSettings *types.SSOClientRedirectSettings
					require.Error(t, sso.ValidateClientRedirect(badURL+"?secret_key=", ceremonyType, defaultSettings))
				})
			}
		}
	})

	t.Run("BadQuery", func(t *testing.T) {
		for _, badURL := range []string{
			"http://127.0.0.1:12345/callback",
			"http://127.0.0.1:12345/callback?secret=a",
			"http://127.0.0.1:12345/callback?secret_key=a&foo=b",
		} {
			for name, ceremonyType := range map[string]sso.CeremonyType{
				"Login": sso.CeremonyTypeLogin,
				"MFA":   sso.CeremonyTypeMFA,
				"Test":  sso.CeremonyTypeTest,
			} {
				t.Run(fmt.Sprintf("Ceremony%v", name), func(t *testing.T) {
					var defaultSettings *types.SSOClientRedirectSettings
					require.Error(t, sso.ValidateClientRedirect(badURL, ceremonyType, defaultSettings))
				})
			}
		}
	})

	t.Run("AllowedHttpsHostnames", func(t *testing.T) {
		for _, goodURL := range []string{
			"https://allowed.domain.invalid/callback",
			"https://foo.allowed.with.subdomain.invalid/callback",
			"https://but.no.subsubdomain.invalid/callback",
		} {
			settings := &types.SSOClientRedirectSettings{
				AllowedHttpsHostnames: []string{
					"allowed.domain.invalid",
					"*.allowed.with.subdomain.invalid",
					"^[-a-zA-Z0-9]+.no.subsubdomain.invalid$",
				},
			}
			for name, ceremonyType := range map[string]sso.CeremonyType{
				"Login": sso.CeremonyTypeLogin,
				"MFA":   sso.CeremonyTypeMFA,
			} {
				t.Run(fmt.Sprintf("Ceremony%v", name), func(t *testing.T) {
					require.NoError(t, sso.ValidateClientRedirect(goodURL+"?secret_key=", ceremonyType, settings))
				})
			}
		}

		for _, badURL := range []string{
			"https://allowed.domain.invalid/notcallback",
			"https://allowed.domain.invalid:12345/callback",
			"http://allowed.domain.invalid/callback",
			"https://not.allowed.domain.invalid/callback",
			"https://allowed.with.subdomain.invalid/callback",
			"https://i.said.no.subsubdomain.invalid/callback",
		} {
			settings := &types.SSOClientRedirectSettings{
				AllowedHttpsHostnames: []string{
					"allowed.domain.invalid",
					"*.allowed.with.subdomain.invalid",
					"^[-a-zA-Z0-9]+.no.subsubdomain.invalid",
				},
			}
			for name, ceremonyType := range map[string]sso.CeremonyType{
				"Login": sso.CeremonyTypeLogin,
				"MFA":   sso.CeremonyTypeMFA,
			} {
				t.Run(fmt.Sprintf("Ceremony%v", name), func(t *testing.T) {
					require.Error(t, sso.ValidateClientRedirect(badURL+"?secret_key=", ceremonyType, settings))
				})
			}
		}
	})

	t.Run("InsecureAllowedCidrRanges", func(t *testing.T) {
		for _, goodURL := range []string{
			"http://192.168.0.27/callback",
			"https://192.168.0.27/callback",
			"http://192.168.0.27:1337/callback",
			"https://192.168.0.27:1337/callback",
			"http://[2001:db8::aaaa:bbbb]/callback",
			"https://[2001:db8::aaaa:bbbb]/callback",
			"http://[2001:db8::aaaa:bbbb]:1337/callback",
			"https://[2001:db8::aaaa:bbbb]:1337/callback",
			"http://[2001:db8::1]/callback",
			"https://[2001:db8::1]/callback",
			"http://[2001:db8::1]:1337/callback",
			"https://[2001:db8::1]:1337/callback",
		} {
			settings := &types.SSOClientRedirectSettings{
				InsecureAllowedCidrRanges: []string{
					"192.168.0.0/24",
					"2001:db8::/96",
				},
			}
			for name, ceremonyType := range map[string]sso.CeremonyType{
				"Login": sso.CeremonyTypeLogin,
				"MFA":   sso.CeremonyTypeMFA,
			} {
				t.Run(fmt.Sprintf("Ceremony%v", name), func(t *testing.T) {
					require.NoError(t, sso.ValidateClientRedirect(goodURL+"?secret_key=", ceremonyType, settings))
				})
			}
		}

		for _, badURL := range []string{
			"http://192.168.1.1/callback",
			"https://192.168.1.1/callback",
			"http://192.168.1.1:80/callback",
			"https://192.168.1.1:443/callback",
			"http://[2001:db8::1:aaaa:bbbb]/callback",
			"https://[2001:db8::1:aaaa:bbbb]/callback",
			"http://[2001:db8::1:aaaa:bbbb]:80/callback",
			"https://[2001:db8::1:aaaa:bbbb]:443/callback",
			"http://[2001:db9::]/callback",
			"https://[2001:db9::]/callback",
			"http://not.an.ip/callback",
			"https://not.an.ip/callback",
			"http://192.168.0.27/nocallback",
			"https://192.168.0.27/nocallback",
			"http://[2001:db8::1]/notacallback",
			"https://[2001:db8::1]/notacallback",
		} {
			settings := &types.SSOClientRedirectSettings{
				InsecureAllowedCidrRanges: []string{
					"192.168.0.0/24",
					"2001:db8::/96",
				},
			}
			for name, ceremonyType := range map[string]sso.CeremonyType{
				"Login": sso.CeremonyTypeLogin,
				"MFA":   sso.CeremonyTypeMFA,
			} {
				t.Run(fmt.Sprintf("Ceremony%v", name), func(t *testing.T) {
					require.Error(t, sso.ValidateClientRedirect(badURL+"?secret_key=", ceremonyType, settings))
				})
			}
		}
	})

	t.Run("SSOTestFlow", func(t *testing.T) {
		for _, goodURL := range []string{
			"http://127.0.0.1:12345/callback",
		} {
			settings := &types.SSOClientRedirectSettings{
				AllowedHttpsHostnames: []string{
					"allowed.domain.invalid",
				},
			}
			require.NoError(t, sso.ValidateClientRedirect(goodURL+"?secret_key=", sso.CeremonyTypeTest, settings))
		}

		for _, badURL := range []string{
			"https://allowed.domain.invalid/callback",
			"http://allowed.domain.invalid/callback",
		} {
			settings := &types.SSOClientRedirectSettings{
				AllowedHttpsHostnames: []string{
					"allowed.domain.invalid",
				},
			}
			require.Error(t, sso.ValidateClientRedirect(badURL+"?secret_key=", sso.CeremonyTypeTest, settings))
		}
	})

	t.Run("SSOMFAWeb", func(t *testing.T) {
		settings := &types.SSOClientRedirectSettings{
			AllowedHttpsHostnames: []string{
				"allowed.domain.invalid",
			},
		}

		// Only allow web mfa redirect as a relative path.
		require.NoError(t, sso.ValidateClientRedirect(sso.WebMFARedirect+"?channel_id=", sso.CeremonyTypeMFA, settings))

		for _, badURL := range []string{
			"localhost:12345" + sso.WebMFARedirect,
			"http://localhost:12345" + sso.WebMFARedirect,
			"https://localhost:12345" + sso.WebMFARedirect,
			"127.0.0.1:12345" + sso.WebMFARedirect,
			"http://127.0.0.1:12345" + sso.WebMFARedirect,
			"https://127.0.0.1:12345" + sso.WebMFARedirect,
			"allowed.domain.invalid" + sso.WebMFARedirect,
			"http://allowed.domain.invalid" + sso.WebMFARedirect,
			"https://allowed.domain.invalid" + sso.WebMFARedirect,
			"not.allowed.domain.invalid" + sso.WebMFARedirect,
			"http://not.allowed.domain.invalid" + sso.WebMFARedirect,
			"https://not.allowed.domain.invalid" + sso.WebMFARedirect,
		} {
			require.Error(t, sso.ValidateClientRedirect(badURL+"?channel_id=", sso.CeremonyTypeMFA, settings))
		}

		// channel_id query parameter must be set.
		require.Error(t, sso.ValidateClientRedirect(sso.WebMFARedirect, sso.CeremonyTypeMFA, settings))

		// Don't allow web mfa redirect for non-mfa ceremonies.
		require.Error(t, sso.ValidateClientRedirect(sso.WebMFARedirect+"?channel_id=", sso.CeremonyTypeLogin, settings))
		require.Error(t, sso.ValidateClientRedirect(sso.WebMFARedirect+"?channel_id=", sso.CeremonyTypeTest, settings))
	})
}

// create a mock proxy server which echos the final proxy redirect destination page. e.g. sso success page.
func newMockProxy(t *testing.T) *httptest.Server {
	mux := http.NewServeMux()
	mux.HandleFunc(sso.LoginSuccessRedirectURL, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(sso.LoginSuccessRedirectURL))
	})
	mux.HandleFunc(sso.LoginFailedRedirectURL, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(sso.LoginFailedRedirectURL))
	})
	mux.HandleFunc(sso.LoginFailedBadCallbackRedirectURL, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(sso.LoginFailedBadCallbackRedirectURL))
	})
	mux.HandleFunc(sso.LoginClose, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(sso.LoginClose))
	})
	mux.HandleFunc(sso.LoginTerminalRedirectURL, func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(sso.LoginTerminalRedirectURL))
	})
	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}
