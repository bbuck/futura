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

import "github.com/spf13/cobra"

// RootCmd is the core command other commands are attached to power the cobra
// CLI library.
var RootCmd = cobra.Command{
	Use:   "futura",
	Short: "Futura is a DB management utility written in Go.",
	Long: `Futura provides a means for managing the state of a database amoung team
members. It's goal is to provide a means for creating/removing/truncating
databases and table as well as managing migrations.`,
}

func init() {
	RootCmd.PersistentFlags().BoolP("dry-run", "D", false, "Run the command (if supported) as a dry run without making change")
}
