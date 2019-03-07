// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package jiri

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

// TestFindRootEnvSymlink checks that FindRoot interprets the value of the
// -root flag as a path, evaluates any symlinks the path might contain, and
// returns the result.
func TestFindRootEnvSymlink(t *testing.T) {
	// Create a temporary directory.
	tmpDir, err := ioutil.TempDir("", "")
	if err != nil {
		t.Fatalf("TempDir() failed: %v", err)
	}
	defer func() { os.RemoveAll(tmpDir) }()

	// Make sure tmpDir is not a symlink itself.
	tmpDir, err = filepath.EvalSymlinks(tmpDir)
	if err != nil {
		t.Fatalf("EvalSymlinks(%v) failed: %v", tmpDir, err)
	}

	// Create a directory and a symlink to it.
	root, perm := filepath.Join(tmpDir, "root"), os.FileMode(0700)
	symRoot := filepath.Join(tmpDir, "sym_root")
	if err := os.MkdirAll(root, perm); err != nil {
		t.Fatalf("%s", err)
	}
	if err := os.Symlink(root, symRoot); err != nil {
		t.Fatalf("%s", err)
	}

	// Set the -root flag to the symlink created above and check that
	// FindRoot() evaluates the symlink.
	rootFlag = symRoot
	if got, want := FindRoot(), root; got != want {
		t.Fatalf("unexpected output: got %v, want %v", got, want)
	}
}
