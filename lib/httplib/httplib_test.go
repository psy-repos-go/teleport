/*
Copyright 2016 Gravitational, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package httplib

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gravitational/roundtrip"
	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/require"
)

type netError struct{}

func (e *netError) Error() string   { return "net" }
func (e *netError) Timeout() bool   { return true }
func (e *netError) Temporary() bool { return true }

func TestConvertResponse(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		err      error
		expected string
	}{
		{
			name: "url error",
			err: &url.Error{
				Op:  "POST",
				URL: "http://localhost",
				Err: errors.New("error goes here"),
			},
			expected: "error goes here",
		},
		{
			name: "url with path error",
			err: &url.Error{
				Op:  "POST",
				URL: "http://localhost?path%20foobar",
				Err: errors.New("error goes here"),
			},
			expected: "error goes here",
		},
		{
			name:     "timeout error",
			err:      &netError{},
			expected: "unable to complete the request due to a timeout, please try again in a few minutes",
		},
		{
			name:     "normal error",
			err:      errors.New("this is a normal error"),
			expected: "this is a normal error",
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := ConvertResponse(&roundtrip.Response{}, test.err)
			require.Error(t, err)
			require.Equal(t, test.expected, err.Error())
		})
	}

}

func TestRewritePaths(t *testing.T) {
	handler := newTestHandler()
	server := httptest.NewServer(
		RewritePaths(handler,
			Rewrite("/v1/sessions/([^/]+)/(.*)", "/v1/namespaces/default/sessions/$1/$2")))
	defer server.Close()
	re, err := http.Post(server.URL+"/v1/sessions/s1/stream", "text/json", nil)
	require.NoError(t, err)
	defer re.Body.Close()
	require.Equal(t, http.StatusOK, re.StatusCode)
	require.Equal(t, "default", handler.capturedNamespace)
	require.Equal(t, "s1", handler.capturedID)

	re, err = http.Post(server.URL+"/v1/namespaces/system/sessions/s2/stream", "text/json", nil)
	require.NoError(t, err)
	defer re.Body.Close()
	require.Equal(t, http.StatusOK, re.StatusCode)
	require.Equal(t, "system", handler.capturedNamespace)
	require.Equal(t, "s2", handler.capturedID)
}

type testHandler struct {
	httprouter.Router
	capturedNamespace string
	capturedID        string
}

func newTestHandler() *testHandler {
	h := &testHandler{}
	h.Router = *httprouter.New()
	h.Router.UseRawPath = true
	h.POST("/v1/sessions/:id/stream", MakeHandler(h.postSessionChunkOriginal))
	h.POST("/v1/namespaces/:namespace/sessions/:id/stream", MakeHandler(h.postSessionChunkNamespace))
	return h
}

func (h *testHandler) postSessionChunkOriginal(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, error) {
	return "ok", nil
}

func (h *testHandler) postSessionChunkNamespace(w http.ResponseWriter, r *http.Request, p httprouter.Params) (interface{}, error) {
	h.capturedNamespace = p.ByName("namespace")
	h.capturedID = p.ByName("id")
	return "ok", nil
}

func TestReadJSON_ContentType(t *testing.T) {
	t.Parallel()

	type TestJSON struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	testCases := []struct {
		name        string
		contentType string
		wantErr     bool
	}{
		{
			name:        "empty value",
			contentType: "",
			wantErr:     true,
		},
		{
			name:        "invalid type",
			contentType: "multipart/form-data",
			wantErr:     true,
		},
		{
			name:        "just type/subtype",
			contentType: "application/json",
		},
		{
			name:        "type/subtype with params",
			contentType: "application/json; charset=utf-8",
		},
	}

	body := TestJSON{Name: "foo", Age: 60}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			payloadBuf := new(bytes.Buffer)
			require.NoError(t, json.NewEncoder(payloadBuf).Encode(body))

			httpReq, err := http.NewRequest("", "", payloadBuf)
			require.NoError(t, err)
			httpReq.Header.Add("Content-Type", tc.contentType)

			output := TestJSON{}
			err = ReadJSON(httpReq, &output)
			if tc.wantErr {
				require.True(t, strings.Contains(err.Error(), "invalid request"))
				require.Empty(t, output)
			} else {
				require.NoError(t, err)
				require.Equal(t, body, output)
			}
		})
	}
}
