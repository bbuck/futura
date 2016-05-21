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
	"sort"

	"github.com/bbuck/futura/adapters"
	"github.com/bbuck/futura/info"
	"github.com/bbuck/futura/log"
	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "adapters",
		Short: "List supported adapters",
		Run: func(_ *cobra.Command, _ []string) {
			keys := make([]string, len(adapters.Map))
			longestKey := -1
			idx := 0
			for key := range adapters.Map {
				keys[idx] = key
				idx++
				if longestKey < 0 || longestKey < len(key) {
					longestKey = len(key)
				}
			}
			sort.Strings(keys)
			fmtString := fmt.Sprintf("%%-%ds - %%s\n", longestKey)
			log.RawPrintf("Supported Adapters -- futura %s\n%s\n", info.Version.String(), ThinSeperator)
			for _, key := range keys {
				adapter := adapters.Map[key]
				log.RawPrintf(fmtString, key, adapter.Description())
			}
			log.NewLine()
		},
	})
}
