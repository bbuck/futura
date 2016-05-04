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

package main

import "github.com/spf13/viper"

func configure() error {
	viper.SetConfigFile("./Futurafile")
	viper.SetConfigType("toml")
	configureEnvs()
	setDefaults()
	err := viper.ReadInConfig()

	return err
}

func configureEnvs() {
	viper.SetEnvPrefix("futura")
	viper.BindEnv("env")
}

func setDefaults() {
	viper.SetDefault("env", "development")
}
