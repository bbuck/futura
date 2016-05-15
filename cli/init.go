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
	"os"

	"github.com/bbuck/futura/config"
	"github.com/spf13/cobra"
)

func init() {
	var (
		adapter string
		force   bool
	)

	initCmd := &cobra.Command{
		Use:   "init",
		Short: "Initialize the directory with a Futurafile for the given adapter.",
		Long: `Initialize the direcotyr with a Futurafile for the given adapter which
provides a jumping off point for the project. If no adapter is provided then
futura will assume a default for you.`,
		Run: func(cmd *cobra.Command, _ []string) {
			_, err := os.Stat("./Futurafile")
			if err != nil && !force {
				fmt.Println("Futurafile already exists, try again with --force if you wish to overwrite it.")
				os.Exit(10)
			} else if err != nil && force {
				os.Remove("./Futurafile")
			}
			file, err := os.Create("./Futurafile")
			if err != nil {
				fmt.Printf("Failed to create initial Futurafile: %s\n", err.Error())
				os.Exit(11)
			}
			defer file.Close()

			config.Futurafile(adapter, file)
		},
	}

	initCmd.Flags().StringVarP(&adapter, "adapter", "a", "pgx", "the adapter with which to generate the Futrafile from")
	initCmd.Flags().BoolVarP(&force, "force", "f", false, "overwrite the Futurafile if one already exists")

	RootCmd.AddCommand(initCmd)
}
