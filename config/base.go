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

package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// SelectedAdapter returns the string identifier for the desired Adapter to use
// for the current environment.
func SelectedAdapter() string {
	return viper.GetString(fmt.Sprintf("%s.adapter", viper.GetString("env")))
}

// EnvSettings returns a map[string]interface{} containing the settings provided
// along with the adapter for the specified environment.
func EnvSettings() map[string]interface{} {
	return viper.GetStringMap(viper.GetString("env"))
}

// DatabaseName returns the database name for the given environment.
func DatabaseName() string {
	return viper.GetString(fmt.Sprintf("%s.database", viper.GetString("env")))
}
