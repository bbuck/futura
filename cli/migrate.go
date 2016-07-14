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
	"log"
	"os"

	"github.com/spf13/cobra"
)

const (
	MigrateDirectionUp   = "up"
	MigrateDirectionDown = "down"
)

func init() {
	migrateCmd := &cobra.Command{
		Aliases: []string{"mig"},
		Use:     "migrate DIRECTION",
		Short:   "Detect and perform directed migrations on the database.",
		Long: `Migrate runs and up or down operation of changes on the database. Up migrations
by default will take any and all migrations that have not been performed.
Optionally a rollback can be performed or a targeted up/down migration for
specific versions.`,
		Run: func(_ *cobra.Command, args []string) {
			if len(args) < 1 {
				log.Println("Please specify whether you want to migrate 'up' or 'down'")
				os.Exit(30)
			} else {
				log.Printf("Direction \"%s\" not implemented", args[0])
			}
		},
	}

	RootCmd.AddCommand(migrateCmd)
}
