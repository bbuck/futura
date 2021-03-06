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
	"os"
	"strconv"
	"strings"

	"github.com/bbuck/futura/config"
	"github.com/bbuck/futura/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

// various values used to generate output
var (
	Seperator     = strings.Repeat("=", 80)
	ThinSeperator = strings.Repeat("-", 80)
)

func getFlag(cmd *cobra.Command, name string) *pflag.Flag {
	flag := cmd.Flags().Lookup(name)
	if flag == nil {
		flag = cmd.PersistentFlags().Lookup(name)
	}

	return flag
}

func getBoolFlag(cmd *cobra.Command, name string) bool {
	flag := getFlag(cmd, name)
	if flag != nil && flag.Value.Type() == "bool" {
		b, _ := strconv.ParseBool(flag.Value.String())

		return b
	}

	return false
}

func validateConfig() {
	err := config.Validate()
	if err != nil {
		log.Errorln(err.Error())
		os.Exit(100)
	}
}
