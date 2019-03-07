// Copyright 2018 The Fuchsia Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"github.com/dahlia-os/jiri"
	"github.com/dahlia-os/jiri/cmdline"
	"github.com/dahlia-os/jiri/project"
)

var resolveFlags struct {
	lockFilePath      string
	localManifestFlag bool
	enablePackageLock bool
	enableProjectLock bool
}

var cmdResolve = &cmdline.Command{
	Runner: jiri.RunnerFunc(runResolve),
	Name:   "resolve",
	Short:  "Generate jiri lockfile",
	Long: `
Generate jiri lockfile in json format for <manifest ...>. If no manifest
provided, jiri will use .jiri_manifest by default.
`,
	ArgsName: "<manifest ...>",
	ArgsLong: "<manifest ...> is a list of manifest files for lockfile generation",
}

func init() {
	flags := &cmdResolve.Flags
	flags.StringVar(&resolveFlags.lockFilePath, "output", "jiri.lock", "Path to the generated lockfile")
	flags.BoolVar(&resolveFlags.localManifestFlag, "local-manifest", false, "Use local manifest")
	flags.BoolVar(&resolveFlags.enablePackageLock, "enable-package-lock", true, "Enable resolving packages in lockfile")
	flags.BoolVar(&resolveFlags.enableProjectLock, "enable-project-lock", false, "Enable resolving projects in lockfile")
}

func runResolve(jirix *jiri.X, args []string) error {
	manifestFiles := make([]string, 0)
	if len(args) == 0 {
		// Use .jiri_manifest if no manifest file path is present
		manifestFiles = append(manifestFiles, jirix.JiriManifestFile())
	} else {
		for _, m := range args {
			manifestFiles = append(manifestFiles, m)
		}
	}
	// While revision pins for projects can be updated by 'jiri edit',
	// instance IDs of packages can only be updated by 'jiri resolve' due
	// to the way how cipd works. Since roller is using 'jiri resolve'
	// to update a single jiri.lock file each time, it will cause conflicting
	// instance ids between updated 'jiri.lock' and un-updated 'jiri.lock' files.
	// Jiri will halt when detecting conflicts in locks. So to make it work,
	// we need to temporarily disable the conflicts detection.
	jirix.IgnoreLockConflicts = true
	return project.GenerateJiriLockFile(jirix, manifestFiles, resolveFlags.lockFilePath, resolveFlags.enableProjectLock, resolveFlags.enablePackageLock, resolveFlags.localManifestFlag)
}
