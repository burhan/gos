/*
 * Copyright 2019 storyicon@foxmail.com
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package install

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/storyicon/gos/pkg/util"
)

// CmdInstall compiles and installs the packages named by the import paths.
var CmdInstall = &cobra.Command{
	Use:   "install [-i] [build flags] [packages]",
	Short: "compile and install packages and dependencies",
	Long: `
Install compiles and installs the packages named by the import paths.

The -i flag installs the dependencies of the named packages as well.

For more about the build flags, see 'go help build'.
For more about specifying packages, see 'go help packages'.

See also: go build, go get, go clean.
`,
	DisableFlagParsing: true,
}

func init() {
	CmdInstall.Run = func(cmd *cobra.Command, args []string) {
		fd := util.GetGoBinaryCMD("install", args)
		fd.Env = util.GetEnvWithLocalProxy()
		fd.Stdout = os.Stdout
		fd.Stderr = os.Stderr
		util.RunCMDWithExit(fd)
	}
}
