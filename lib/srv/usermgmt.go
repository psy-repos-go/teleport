/*
Copyright 2022 Gravitational, Inc.

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

package srv

import (
	"os/exec"
	"os/user"

	"github.com/gravitational/teleport/api/types"
	"github.com/gravitational/trace"
	"github.com/siddontang/go/log"
)

// todo(amk): make group/user management an interface, or make the
//            operations overridable somehow for testing

const groupExistExit = 9 // man GROUPADD(8), exit codes section
const userExistExit = 9  // man USERADD(8), exit codes section

func groupAdd(groupname string) (exitCode int, err error) {
	groupaddBin, err := exec.LookPath("groupadd")
	if err != nil {
		return -1, trace.Wrap(err, "cant find groupadd binary")
	}
	cmd := exec.Command(groupaddBin, groupname)
	err = cmd.Run()
	return cmd.ProcessState.ExitCode(), err
}

func userAdd(username string, groups []string) (exitCode int, err error) {
	useraddBin, err := exec.LookPath("useradd")
	if err != nil {
		return -1, trace.Wrap(err, "cant find useradd binary")
	}
	// useradd --create-home (username) (groups)...
	args := append([]string{"--create-home", username}, groups...)
	cmd := exec.Command(useraddBin, args...)
	err = cmd.Run()
	return cmd.ProcessState.ExitCode(), err
}

func addUserToGroups(username string, groups []string) (exitCode int, err error) {
	usermodBin, err := exec.LookPath("usermod")
	if err != nil {
		return -1, trace.Wrap(err, "cant find usermod binary")
	}
	args := []string{"-aG"}
	args = append(args, groups...)
	args = append(args, username)
	// usermod -aG (append groups) (username)
	cmd := exec.Command(usermodBin, args...)
	err = cmd.Run()
	return cmd.ProcessState.ExitCode(), err
}

func userDel(username string) (exitCode int, err error) {
	userdelBin, err := exec.LookPath("userdel")
	if err != nil {
		return -1, trace.Wrap(err, "cant find userdel binary")
	}
	// userdel -r (remove home) username
	cmd := exec.Command(userdelBin, "-r", username)
	err = cmd.Run()
	return cmd.ProcessState.ExitCode(), err
}

type userCloser struct {
	user string
}

func (u *userCloser) deleteUserInTeleportGroup() error {
	tempUser, err := user.Lookup(u.user)
	if err != nil {
		return trace.Wrap(err)
	}
	ids, err := tempUser.GroupIds()
	if err != nil {
		return trace.Wrap(err)
	}
	teleportGroup, err := user.LookupGroup(types.TeleportServiceGroup)
	for _, id := range ids {
		if id == teleportGroup.Gid {
			_, err := userDel(u.user)
			return trace.Wrap(err)
		}
	}
	log.Debug("Not deleting user %q, not a temporary user", u.user)
	return nil
}

func (u *userCloser) Close() error {
	return trace.Wrap(u.deleteUserInTeleportGroup())
}

func createTeleportServiceGroupIfNotExist() error {
	_, err := user.LookupGroup(types.TeleportServiceGroup)
	if err != nil && err == user.UnknownGroupError(types.TeleportServiceGroup) {
		_, err := groupAdd(types.TeleportServiceGroup)
		return trace.Wrap(err)
	}
	return trace.Wrap(err)
}

func createTemporaryUserAndAddToTeleportGroup(username string, groups []string) (*userCloser, bool, error) {
	err := createTeleportServiceGroupIfNotExist()
	if err != nil {
		return nil, false, trace.Wrap(err)
	}
	tempUser, err := user.Lookup(username)
	if tempUser != nil {
		return nil, false, trace.AlreadyExists("User already exists")
	}
	if err != nil && err != user.UnknownUserError(username) {
		return nil, false, trace.Wrap(err)
	}
	// TODO(lxea): create groups
	code, err := userAdd(username, groups)
	if code != userExistExit && err != nil {
		return nil, false, trace.Wrap(err)
	}
	_, err = addUserToGroups(username, []string{types.TeleportServiceGroup})
	if err != nil {
		return nil, false, err
	}
	return &userCloser{
		user: username,
	}, true, nil
}
