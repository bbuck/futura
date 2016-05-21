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
	"io"
	"text/template"

	"github.com/bbuck/futura/adapters"
	"github.com/bbuck/futura/info"
)

const futurafile = `# Default config file generated by futura {{ .Version }}

# The Futurafile is segmented by environment names. The commonly used names
# are 'development,' 'qa,' and 'production.' Feel free to add any additional
# environment names as you see fit. This allows you to control you configuration
# by different environments.
#
# The minimum requirement for an environment configuration is the 'adapter.'
# This is what instructs futura how to connect to your database. Each adapter
# requires in own configurations.

[development]

{{ .AdapterDetails }}

[qa]

{{ .AdapterDetails }}

[production]

{{ .AdapterDetails }}
`

// Futurafile returns a basic generated Futurafile for use
func Futurafile(adapter string, w io.Writer) {
	var details string
	if adapter, ok := adapters.Map[adapter]; ok {
		if d, ok := adapter.(adapters.Details); ok {
			details = d.DefaultConfigString()
		}
	}
	data := struct {
		Version        string
		AdapterDetails string
	}{
		Version:        info.Version.String(),
		AdapterDetails: details,
	}
	t := template.Must(template.New("").Parse(futurafile))
	t.Execute(w, data)
}
