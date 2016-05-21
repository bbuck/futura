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
	"path/filepath"
	"text/template"
	"time"

	"github.com/bbuck/futura/adapters"
	"github.com/bbuck/futura/config"
	"github.com/bbuck/futura/log"
	"github.com/spf13/cobra"
)

var migrationTemplate = `{{ .UpComment }}
{{ .CommentStart }} perform your up migration here

{{ .DownComment }}
{{ .CommentStart }} perform your down migration here
`

var migFolderPerm os.FileMode = 0755

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Aliases: []string{"gen"},
		Use:     "generate NAME",
		Short:   "Generate a migration file for the current adapter",
		Long: `Given a name for a file, this will generate a new file in the migration
path with a timestamped version value, ready for migrations to be made.`,
		Run: func(_ *cobra.Command, args []string) {
			validateConfig()

			if len(args) == 0 {
				log.Errorf("You must provide a name for the new migration file.")
				os.Exit(20)
			}

			if adapter, ok := adapters.Map[config.SelectedAdapter()]; ok {
				filename := args[0]
				_, err := os.Stat(config.MigrationsPath())
				if err != nil && os.IsNotExist(err) {
					os.MkdirAll(config.MigrationsPath(), migFolderPerm)
				}
				filename = fmt.Sprintf("%d.%s.%s", time.Now().UnixNano(), filename, adapter.RawFileExtension())

				file, err := os.Create(filepath.Join(config.MigrationsPath(), filename))
				if err != nil {
					log.Errorf(err.Error())
					os.Exit(22)
				}
				defer file.Close()

				tmpl := template.New("")
				tmpl = template.Must(tmpl.Parse(migrationTemplate))
				tmpl.Execute(file, adapter)

				log.Printf("Migration %q has been successfully created!", filename)
			} else {
				log.Errorf("The chosen adapter %q does not exist", config.SelectedAdapter())
				os.Exit(21)
			}
		},
	})
}
