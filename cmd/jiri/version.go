// Copyright 2016 The Fuchsia Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"

	"github.com/dahlia-os/jiri/cmdline"
	"github.com/dahlia-os/jiri/version"
)

var cmdVersion = &cmdline.Command{
	Runner: cmdline.RunnerFunc(runVersion),
	Name:   "version",
	Short:  "Print the jiri version",
	Long: `
Print the Git commit revision jiri was built from and the build date.
`,
}

func runVersion(env *cmdline.Env, args []string) error {
	var versionString bytes.Buffer
	fmt.Fprintf(&versionString, "Jiri")

	v := version.FormattedVersion()
	if v != "" {
		fmt.Fprintf(&versionString, " %s", v)
	}

	fmt.Printf("%s\n", versionString.String())

	return nil
}
