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

package log

import (
	"fmt"
	"os"
)

// Println uses Printf for each item passed to it, assuming a format string of
// "%s"
func Println(items ...interface{}) {
	for _, item := range items {
		Printf("%s", item)
	}
}

// Printf takes a format string, appends two newlines to the end of the string
// and then prints it to the console. In the future this may be more useful,
// but for now it's to keep my laziness happy.
func Printf(format string, items ...interface{}) {
	format = fmt.Sprintf("%s\n\n", format)
	RawPrintf(format, items...)
}

// RawPrintf is here when you want absolute control over the output, no additional
// notations (like the two newlines)
func RawPrintf(format string, items ...interface{}) {
	fmt.Printf(format, items...)
}

// Errorln uses Errorf for each item passed to it, assuming a format string of
// "%s"
func Errorln(items ...interface{}) {
	for _, item := range items {
		Errorf("%s", item)
	}
}

// Errorf is like Printf except that it prepends 'ERROR: ' and logs to stderr
func Errorf(format string, items ...interface{}) {
	format = fmt.Sprintf("ERROR: %s\n\n", format)
	fmt.Fprintf(os.Stderr, format, items...)
}
