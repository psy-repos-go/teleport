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

package main

import (
	"os"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestLoadConfigNonExistingFile(t *testing.T) {
	fullFilePath := "/tmp/doesntexist." + uuid.NewString()
	gotConfig, gotErr := loadConfig(fullFilePath)
	require.NoError(t, gotErr)
	require.Equal(t, &TshConfig{}, gotConfig)
}

func TestLoadConfigEmptyFile(t *testing.T) {
	file, err := os.CreateTemp("", "test-telelport")
	require.NoError(t, err)
	defer os.Remove(file.Name())

	_, err = file.Write([]byte(" "))
	require.NoError(t, err)

	gotConfig, gotErr := loadConfig(file.Name())
	require.NoError(t, gotErr)
	require.Equal(t, &TshConfig{}, gotConfig)
}
