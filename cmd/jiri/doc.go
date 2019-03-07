// Copyright 2015 The Vanadium Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file was auto-generated via go generate.
// DO NOT UPDATE MANUALLY

/*
Command jiri is a multi-purpose tool for multi-repo development.

Usage:
   jiri [flags] <command>

The jiri commands are:
   cl          Manage changelists for multiple projects
   import      Adds imports to .jiri_manifest file
   project     Manage the jiri projects
   snapshot    Manage project snapshots
   update      Update all jiri projects
   which       Show path to the jiri tool
   runp        Run a command in parallel across jiri projects
   help        Display help for commands or topics

The jiri additional help topics are:
   filesystem  Description of jiri file system layout
   manifest    Description of manifest files

The jiri flags are:
 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

The global flags are:
 -metadata=<just specify -metadata to activate>
   Displays metadata for the program and exits.
 -time=false
   Dump timing information to stderr before exiting the program.

Jiri cl - Manage changelists for multiple projects

Manage changelists for multiple projects.

Usage:
   jiri cl [flags] <command>

The jiri cl commands are:
   cleanup     Clean up changelists that have been merged
   upload      Upload a changelist for review
   new         Create a new local branch for a changelist
   sync        Bring a changelist up to date

The jiri cl flags are:
 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri cl cleanup - Clean up changelists that have been merged

Command "cleanup" checks that the given branches have been merged into the
corresponding remote branch. If a branch differs from the corresponding remote
branch, the command reports the difference and stops. Otherwise, it deletes the
given branches.

Usage:
   jiri cl cleanup [flags] <branches>

<branches> is a list of branches to cleanup.

The jiri cl cleanup flags are:
 -f=false
   Ignore unmerged changes.
 -remote-branch=master
   Name of the remote branch the CL pertains to, without the leading "origin/".

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri cl upload - Upload a changelist for review

Command "upload" squashes all commits of a local branch into a single "changelist"
and uploads this changelist to Gerrit as a single commit. First time the command
is invoked, it generates a Change-Id for the changelist, which is appended to
the commit message. Consecutive invocations of the command use the same
Change-Id by default, informing Gerrit that the incomming commit is an update of
an existing changelist.

Usage:
   jiri cl upload [flags]

The jiri cl upload flags are:
 -autosubmit=false
   Automatically submit the changelist when feasible.
 -cc=
   Comma-seperated list of emails or LDAPs to cc.
 -check-uncommitted=true
   Check that no uncommitted changes exist.
 -clean-multipart-metadata=false
   Cleanup the metadata associated with multipart CLs pertaining the MultiPart:
   x/y message without uploading any CLs.
 -commit-message-body-file=
   file containing the body of the CL description, that is, text without a
   ChangeID, MultiPart etc.
 -current-project-only=false
   Run upload in the current project only.
 -d=false
   Send a draft changelist.
 -edit=true
   Open an editor to edit the CL description.
 -host=
   Gerrit host to use.  Defaults to gerrit host specified in manifest.
 -m=
   CL description.
 -presubmit=all
   The type of presubmit tests to run. Valid values: none,all.
 -r=
   Comma-seperated list of emails or LDAPs to request review.
 -remote-branch=master
   Name of the remote branch the CL pertains to, without the leading "origin/".
 -set-topic=true
   Set Gerrit CL topic.
 -topic=
   CL topic, defaults to <username>-<branchname>.
 -verify=true
   Run pre-push git hooks.

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri cl new - Create a new local branch for a changelist

Command "new" creates a new local branch for a changelist. In particular, it
forks a new branch with the given name from the current branch and records the
relationship between the current branch and the new branch in the .jiri metadata
directory. The information recorded in the .jiri metadata directory tracks
dependencies between CLs and is used by the "jiri cl sync" and "jiri cl upload"
commands.

Usage:
   jiri cl new [flags] <name>

<name> is the changelist name.

The jiri cl new flags are:
 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri cl sync - Bring a changelist up to date

Command "sync" brings the CL identified by the current branch up to date with
the branch tracking the remote branch this CL pertains to. To do that, the
command uses the information recorded in the .jiri metadata directory to
identify the sequence of dependent CLs leading to the current branch. The
command then iterates over this sequence bringing each of the CLs up to date
with its ancestor. The end result of this process is that all CLs in the
sequence are up to date with the branch that tracks the remote branch this CL
pertains to.

NOTE: It is possible that the command cannot automatically merge changes in an
ancestor into its dependent. When that occurs, the command is aborted and prints
instructions that need to be followed before the command can be retried.

Usage:
   jiri cl sync [flags]

The jiri cl sync flags are:
 -remote-branch=master
   Name of the remote branch the CL pertains to, without the leading "origin/".

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri import

Command "import" adds imports to the [root]/.jiri_manifest file, which specifies
manifest information for the jiri tool.  The file is created if it doesn't
already exist, otherwise additional imports are added to the existing file.

An <import> element is added to the manifest representing a remote manifest
import.  The manifest file path is relative to the root directory of the remote
import repository.

Example:
  $ jiri import myfile https://foo.com/bar.git

Run "jiri help manifest" for details on manifests.

Usage:
   jiri import [flags] <manifest> <remote>

<manifest> specifies the manifest file to use.

<remote> specifies the remote manifest repository.

The jiri import flags are:
 -name=manifest
   The name of the remote manifest project.
 -out=
   The output file.  Uses [root]/.jiri_manifest if unspecified.  Uses stdout
   if set to "-".
 -overwrite=false
   Write a new .jiri_manifest file with the given specification.  If it already
   exists, the existing content will be ignored and the file will be
   overwritten.
 -protocol=git
   The version control protocol used by the remote manifest project.
 -remote-branch=master
   The branch of the remote manifest project to track, without the leading
   "origin/".
 -root=
   Root to store the manifest project locally.

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri patch - Patch in the existing change

Command "patch" applies the existing changelist to the current project. The
change can be identified either using change ID, in which case the latest
patchset will be used, or the the full reference.

A new branch will be created to apply the patch to. The default name of this
branch is "change/<changeset>/<patchset>", but this can be overriden using the
-branch flag. The command will fail if the branch already exists. The -delete
flag will delete the branch if already exists. Use the -force flag to force
deleting the branch even if it contains unmerged changes).

Usage:
   jiri patch <change>

<change> is a change ID or a full reference.

The jiri project info flags are:
 -branch=
   Name of the branch the patch will be applied to
 -delete=false
   Delete the existing branch if already exists
 -force=false
   Use force when deleting the existing branch
 -host=
   Gerrit host to use.  Defaults to gerrit host specified in manifest.

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri project - Manage the jiri projects

Manage the jiri projects.

Usage:
   jiri project [flags] <command>

The jiri project commands are:
   clean        Restore jiri projects to their pristine state
   info         Provided structured input for existing jiri projects and
                branches
   list         List existing jiri projects and branches
   shell-prompt Print a succinct status of projects suitable for shell prompts

The jiri project flags are:
 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri project clean - Restore jiri projects to their pristine state

Restore jiri projects back to their master branches and get rid of all the local
branches and changes.

Usage:
   jiri project clean [flags] <project ...>

<project ...> is a list of projects to clean up.

The jiri project clean flags are:
 -branches=false
   Delete all non-master branches.

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri project info - Provided structured input for existing jiri projects and branches

Inspect the local filesystem and provide structured info on the existing
projects and branches. Projects are specified using either names or regular
expressions that are matched against project names. If no command line
arguments are provided the project that the contains the current directory is
used, or if run from outside of a given project, all projects will be used. The
information to be displayed can be specified using a Go template, supplied via
the -template flag.

Usage:
   jiri project info [flags] <project-names>...

<project-names>... a list of project names

The jiri project info flags are:
 -json-output=
   Path to write operation results to.
 -regexp=false
   Use argument as regular expression.
 -template=
   The template for the fields to display.

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri project list - List existing jiri projects and branches

Inspect the local filesystem and list the existing projects and branches.

Usage:
   jiri project list [flags]

The jiri project list flags are:
 -branches=false
   Show project branches.
 -nopristine=false
   If true, omit pristine projects, i.e. projects with a clean master branch and
   no other branches.

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri project shell-prompt - Print a succinct status of projects suitable for shell prompts

Reports current branches of jiri projects (repositories) as well as an
indication of each project's status:
  *  indicates that a repository contains uncommitted changes
  %  indicates that a repository contains untracked files

Usage:
   jiri project shell-prompt [flags]

The jiri project shell-prompt flags are:
 -check-dirty=true
   If false, don't check for uncommitted changes or untracked files. Setting
   this option to false is dangerous: dirty master branches will not appear in
   the output.
 -show-name=false
   Show the name of the current repo.

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri snapshot - Create a new project snapshot

The "jiri snapshot <snapshot>" command captures the current project state
in a manifest.

Usage:
   jiri snapshot [flags] <snapshot>

<snapshot> is the snapshot manifest file.

The jiri snapshot create flags are:
 -time-format=2006-01-02T15:04:05Z07:00
   Time format for snapshot file name.

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri update - Update all jiri projects

Updates all projects. The sequence in which the individual updates happen
guarantees that we end up with a consistent workspace. The set of projects
to update is described in the manifest.

Run "jiri help manifest" for details on manifests.

Usage:
   jiri update [flags] <snapshot>

<snapshot> is the snapshot manifest file.

The jiri update flags are:
 -attempts=1
   Number of attempts before failing.
 -gc=false
   Garbage collect obsolete repositories.
 -manifest=
   Name of the project manifest.

 -color=true
   Use color to format output.
 -v=false
   Print verbose output.

Jiri runp - Run a command in parallel across jiri projects

Run a command in parallel across one or more jiri projects. Commands are run
using the shell specified by the users $SHELL environment variable, or "sh"
if that's not set. Thus commands are run as $SHELL -c "args..."

Usage:
   jiri runp [flags] <command line>

A command line to be run in each project specified by the supplied command line
flags. Any environment variables intended to be evaluated when the command line
is run must be quoted to avoid expansion before being passed to runp by the
shell.

The jiri runp flags are:
 -collate-stdout=true
   Collate all stdout output from each parallel invocation and display it as if
   had been generated sequentially. This flag cannot be used with
   -show-name-prefix, -show-key-prefix or -interactive.
 -env=
   specify an environment variable in the form: <var>=[<val>],...
 -exit-on-error=false
   If set, all commands will killed as soon as one reports an error, otherwise,
   each will run to completion.
 -has-branch=
   A regular expression specifying branch names to use in matching projects. A
   project will match if the specified branch exists, even if it is not checked
   out.
 -has-gerrit-message=false
   If specified, match branches that have, or have no, gerrit message
 -has-uncommitted=false
   If specified, match projects that have, or have no, uncommitted changes
 -has-untracked=false
   If specified, match projects that have, or have no, untracked files
 -interactive=true
   If set, the command to be run is interactive and should not have its
   stdout/stderr manipulated. This flag cannot be used with -show-name-prefix,
   -show-key-prefix or -collate-stdout.
 -merge-policies=+CCFLAGS,+CGO_CFLAGS,+CGO_CXXFLAGS,+CGO_LDFLAGS,+CXXFLAGS,GOARCH,GOOS,GOPATH:,^GOROOT*,+LDFLAGS,:PATH,VDLPATH:
   specify policies for merging environment variables
 -projects=
   A Regular expression specifying project keys to run commands in. By default,
   runp will use projects that have the same branch checked as the current
   project unless it is run from outside of a project in which case it will
   default to using all projects.
 -show-key-prefix=false
   If set, each line of output from each project will begin with the key of the
   project followed by a colon. This is intended for use with long running
   commands where the output needs to be streamed. Stdout and stderr are spliced
   apart. This flag cannot be used with -interactive, -show-name-prefix or
   -collate-stdout
 -show-name-prefix=false
   If set, each line of output from each project will begin with the name of the
   project followed by a colon. This is intended for use with long running
   commands where the output needs to be streamed. Stdout and stderr are spliced
   apart. This flag cannot be used with -interactive, -show-key-prefix or
   -collate-stdout.
 -v=false
   Print verbose logging information

 -color=true
   Use color to format output.

Jiri help - Display help for commands or topics

Help with no args displays the usage of the parent command.

Help with args displays the usage of the specified sub-command or help topic.

"help ..." recursively displays help for all commands and topics.

Usage:
   jiri help [flags] [command/topic ...]

[command/topic ...] optionally identifies a specific sub-command or help topic.

The jiri help flags are:
 -style=compact
   The formatting style for help output:
      compact   - Good for compact cmdline output.
      full      - Good for cmdline output, shows all global flags.
      godoc     - Good for godoc processing.
      shortonly - Only output short description.
   Override the default by setting the CMDLINE_STYLE environment variable.
 -width=<terminal width>
   Format output to this target width in runes, or unlimited if width < 0.
   Defaults to the terminal width if available.  Override the default by setting
   the CMDLINE_WIDTH environment variable.

Jiri filesystem - Description of jiri file system layout

All data managed by the jiri tool is located in the file system under a root
directory, colloquially called the jiri root directory.  The file system layout
looks like this:

 [root]                              # root directory (name picked by user)
 [root]/.jiri_root                   # root metadata directory
 [root]/.jiri_root/bin               # contains tool binaries (jiri, etc.)
 [root]/.jiri_root/update_history    # contains history of update snapshots
 [root]/.manifest                    # contains jiri manifests
 [root]/[project1]                   # project directory (name picked by user)
 [root]/[project1]/.jiri             # project metadata directory
 [root]/[project1]/.jiri/metadata.v2 # project metadata file
 [root]/[project1]/.jiri/<<cls>>     # project per-cl metadata directories
 [root]/[project1]/<<files>>         # project files
 [root]/[project2]...

The [root] and [projectN] directory names are picked by the user.  The <<cls>>
are named via jiri cl new, and the <<files>> are named as the user adds files
and directories to their project.  All other names above have special meaning to
the jiri tool, and cannot be changed; you must ensure your path names don't
collide with these special names.

To find the [root] directory, the jiri binary looks for the .jiri_root
directory, starting in the current working directory and walking up the
directory chain. The search is terminated successfully when the .jiri_root
directory is found; it fails after it reaches the root of the file system.
Thus jiri must be invoked from the [root] directory or one of its
subdirectories.  To invoke jiri from a different directory, you can set the
-root flag to point to your [root] directory.

Keep in mind that when "jiri update" is run, the jiri tool itself is
automatically updated along with all projects.  Note that if you have multiple
[root] directories on your file system, you must remember to run the jiri
binary corresponding to your [root] directory.  Things may fail if you mix
things up, since the jiri binary is updated with each call to "jiri update",
and you may encounter version mismatches between the jiri binary and the
various metadata files or other logic.

The jiri binary is located at [root]/.jiri_root/bin/jiri

Jiri manifest - Description of manifest files

Jiri manifest files describe the set of projects that get synced when running
"jiri update".

The first manifest file that jiri reads is in [root]/.jiri_manifest.  This
manifest **must** exist for the jiri tool to work.

Usually the manifest in [root]/.jiri_manifest will import other manifests from
remote repositories via <import> tags, but it can contain its own list of
projects as well.

Manifests have the following XML schema:

<manifest>
  <imports>
    <import remote="https://vanadium.googlesource.com/manifest"
            manifest="public"
            name="manifest"
    />
    <localimport file="/path/to/local/manifest"/>
    ...
  </imports>
  <projects>
    <project name="my-project"
             path="path/where/project/lives"
             protocol="git"
             remote="https://github.com/myorg/foo"
             revision="ed42c05d8688ab23"
             remotebranch="my-branch"
             gerrithost="https://myorg-review.googlesource.com"
             githooks="path/to/githooks-dir"
    />
    ...
  </projects>
  <hooks>
    <hook name="update"
          project="mojo/public"
          action="update.sh"/>
    ...
  </hooks>

</manifest>

The <import> and <localimport> tags can be used to share common projects across
multiple manifests.

A <localimport> tag should be used when the manifest being imported and the
importing manifest are both in the same repository, or when neither one is in a
repository.  The "file" attribute is the path to the manifest file being
imported.  It can be absolute, or relative to the importing manifest file.

If the manifest being imported and the importing manifest are in different
repositories then an <import> tag must be used, with the following attributes:

* remote (required) - The remote url of the repository containing the manifest
to be imported

* manifest (required) - The path of the manifest file to be imported, relative
to the repository root.

* name (optional) - The name of the project corresponding to the manifest
repository.  If your manifest contains a <project> with the same remote as the
manifest remote, then the "name" attribute of on the <import> tag should match
the "name" attribute on the <project>.  Otherwise, jiri will clone the manifest
repository on every update.

The <project> tags describe the projects to sync, and what state they should
sync to, accoring to the following attributes:

* name (required) - The name of the project.

* path (required) - The location where the project will be located, relative to
the jiri root.

* remote (required) - The remote url of the project repository.

* protocol (optional) - The protocol to use when cloning and syncing the repo.
Currently "git" is the default and only supported protocol.

* remotebranch (optional) - The remote branch that the project will sync to.
Defaults to "master".  The "remotebranch" attribute is ignored if "revision" is
specified.

* revision (optional) - The specific revision (usually a git SHA) that the
project will sync to.  If "revision" is  specified then the "remotebranch"
attribute is ignored.

* gerrithost (optional) - The url of the Gerrit host for the project.  If
specified, then running "jiri cl upload" will upload a CL to this Gerrit host.

* githooks (optional) - The path (relative to [root]) of a directory containing
git hooks that will be installed in the projects .git/hooks directory during
each update.

The <hook> tag describes the hooks that must be executed after every 'jiri update'
They are configured via the following attributes:

* name (required) - The name of the of the hook to identify it

* project (required) - The name of the project where the hook is present

* action (required) - Action to be performed inside the project.
It is mostly identified by a script
*/
package main
