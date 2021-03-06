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

package info

import "fmt"

// Version is the cononical version string for the future application.
var Version = V{
	major: 0,
	minor: 1,
	build: 0,
	tag:   "alpha",
}

// V represents a version for the application.
type V struct {
	major, minor, build int
	tag                 string
}

// String converts the version into a readable string that can be displayed to
// the user.
func (v V) String() string {
	str := fmt.Sprintf("%d.%d.%d", v.major, v.minor, v.build)
	if len(v.tag) > 0 {
		str += fmt.Sprintf(" %s", v.tag)
	}

	return str
}
