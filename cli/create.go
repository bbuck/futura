// Copyright 2016 Brandon Buck
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cli

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create the database for the specified environment",
		Long: `Run the creation command to create the database for the specified
environment. If you would like to preview the query that will happen you can
specify this as a dry run.`,
		Run: func(cmd *cobra.Command, _ []string) {
			dryRun := getBoolFlag(cmd, "dry-run")
			if !dryRun {
				fmt.Println(viper.GetString("env"))
			}
		},
	}

	RootCmd.AddCommand(cmd)
}
