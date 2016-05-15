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

import (
	"fmt"

	"github.com/spf13/viper"
)

// NoAdapterError provides a clean way to define which environment did not
// specify an adapter. Adapters are required in order to load in the the
// functionality required to modify the database type for the selected
// environment.
type NoAdapterError string

// Error returns a string defining the environment that doesn't define an
// adapter.
func (n NoAdapterError) Error() string {
	return fmt.Sprintf("There is no adapter configured for the environment %q", string(n))
}

// C is the config values for this project, which is just an exported Viper
// instance.
var C *viper.Viper

func init() {
	C = viper.New()
}

// Configure sets up the information necessary to load configuration data for
// the library.
func Configure() error {
	C.SetConfigFile("./Futurafile")
	C.SetConfigType("toml")
	configureEnvs()
	configureFlags()
	setDefaults()
	if err := C.ReadInConfig(); err != nil {
		return err
	}

	return validateConfig()
}

// check configuration data for required fields
func validateConfig() error {
	env := C.GetString("env")
	if !C.IsSet(fmt.Sprintf("%s.adapter", env)) {
		return NoAdapterError(env)
	}

	return nil
}

// add evnironment links to the config object
func configureEnvs() {
	C.SetEnvPrefix("futura")
	C.BindEnv("env")
}

// add flags to the config object
func configureFlags() {
	// empty for now
}

// set default values
func setDefaults() {
	C.SetDefault("env", "development")
}
