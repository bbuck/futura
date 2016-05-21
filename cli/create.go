// Copyright 2016 Futura Contributors
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
	"log"
	"os"

	"github.com/bbuck/futura/adapters"
	"github.com/bbuck/futura/config"
	"github.com/spf13/cobra"
)

func init() {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "Create the database for the specified environment",
		Long: `Run the creation command to create the database for the specified
environment. If you would like to preview the query that will happen you can
specify this as a dry run.`,
		Run: func(cmd *cobra.Command, _ []string) {
			validateConfig()

			dbName := config.DatabaseName()
			if len(dbName) == 0 {
				fmt.Fprintf(os.Stderr, "No database given to create!\n")
				os.Exit(5)
			}
			log.Printf("%s\ncreating database %q ...\n%s", Seperator, dbName, Seperator)

			dryRun := getBoolFlag(cmd, "dry-run")
			if !dryRun {
				if adapter, ok := adapters.Map[config.SelectedAdapter()]; ok {
					err := adapter.CreateDatabase(config.DatabaseName())
					if err != nil {
						fmt.Fprintf(os.Stderr, "Failed to create database: %#v", err)
						os.Exit(7)
					}
				} else {
					fmt.Fprintf(os.Stderr, "There is no adapter defined for %q", config.SelectedAdapter())
					os.Exit(6)
				}
			}
		},
	}

	RootCmd.AddCommand(cmd)
}
