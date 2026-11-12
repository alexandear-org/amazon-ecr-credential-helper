// Copyright Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package config

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandHome(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("USERPROFILE", home)

	tests := []struct {
		path string
		want string
	}{
		{path: "", want: ""},
		{path: "/foo/bar", want: "/foo/bar"},
		{path: "~", want: home},
		{path: "~/.ecr", want: filepath.Join(home, ".ecr")},
		{path: "~/.ecr/log", want: filepath.Join(home, ".ecr", "log")},
	}
	for _, tt := range tests {
		got, err := ExpandHome(tt.path)
		assert.NoError(t, err, tt.path)
		assert.Equal(t, tt.want, got, tt.path)
	}
}

func TestExpandHomeOtherUser(t *testing.T) {
	_, err := ExpandHome("~other/.ecr")
	assert.Error(t, err)
}
