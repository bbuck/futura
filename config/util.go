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

package config

import "bytes"

// SelectedAdapter returns the string identifier for the desired Adapter to use
// for the current environment.
func SelectedAdapter() string {
	return C.GetString(EnvKey("adapter"))
}

// EnvSettings returns a map[string]interface{} containing the settings provided
// along with the adapter for the specified environment.
func EnvSettings() map[string]interface{} {
	return C.GetStringMap(env())
}

// DatabaseName returns the database name for the given environment.
func DatabaseName() string {
	return C.GetString(EnvKey("database"))
}

// EnvKey returns a new dot seperate key prepended with the environment.
func EnvKey(keys ...string) string {
	buf := bytes.NewBufferString(env())
	buf.WriteRune('.')
	for _, key := range keys[:len(keys)-1] {
		buf.WriteString(key)
		buf.WriteRune('.')
	}
	buf.WriteString(keys[len(keys)-1])

	return buf.String()
}

// cache the env value for use in retrieving config values
var envCache string

func env() string {
	if len(envCache) == 0 {
		envCache = C.GetString("env")
	}

	return envCache
}
