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

package adapters

import (
	"io/ioutil"
	"strings"
)

// Map will map the user facing adpater names to an Adapter object that
// will begin the process of powering the futura migration tool.
var Map = make(map[string]A)

// MigrationDirection defines a type to represent the direction of a migration
type MigrationDirection bool

// Migration directions
const (
	MigrateUp   = MigrationDirection(true)
	MigrateDown = MigrationDirection(false)
)

// VersionTableName is the default name for the versions table (if necessary)
// to track which migrations have been executed. The specifics of this table
// are to be defined by the adapter if needed. This name is optional too, although
// highly recommended.
const VersionTableName = "futura_versions"

// An A defines a base interface for entrypoints. This is bare minimum to
// enable an Adapter to work with Futura via raw scripts.
type A interface {
	CreateDatabase(string) error
	CreateVersionTable() error
	RunRawMigration(string) error
	RawFileExtension() string
	RawFileParser(string) (*FileParser, error)
	UpComment() string
	DownComment() string
	CommentStart() string
	DefaultConfigString() string
	Description() string
}

// FileParser will break apart a raw query file to extract the up and down
// queries written in the adapters native query language.
type FileParser struct {
	Filename    string
	upComment   string
	downComment string
	file        string
}

// NewFileParser creates a file parser that can extract the up migration and
// down migrations from the given file based on information provided by the
// Adapter.
func NewFileParser(filename, upComment, downComment string) (*FileParser, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return &FileParser{
		Filename:    filename,
		upComment:   upComment,
		downComment: downComment,
		file:        string(contents),
	}, nil
}

// UpQuery extracts and returns the Up migration from the migration file.
func (f *FileParser) UpQuery() string {
	return queryFromString(f.file, f.upComment, f.downComment)
}

// DownQuery extracts and returns the Down migration from the migration file.
func (f *FileParser) DownQuery() string {
	return queryFromString(f.file, f.downComment, f.upComment)
}

// queryFromString extracts the query from the string from the start to end
// values given. If end is before start, the end of file is assumed.
func queryFromString(file, start, end string) string {
	startIdx := strings.Index(file, start)
	if startIdx < 0 {
		return ""
	}
	endIdx := strings.Index(file, end)
	if endIdx < 0 || endIdx < startIdx {
		endIdx = len(file)
	}
	// NOTE: Shouldn't happen, but let's go ahead and be pre-emptive about
	// catching it anyways.
	if startIdx == endIdx {
		return ""
	}

	return strings.TrimSpace(file[startIdx:endIdx])
}

// type ColumnConfig struct {
// 	Name          string
// 	Type          string
// 	Default       bool
// 	PrimaryKey    bool
// 	AutoIncrement bool
// 	Key           bool
// 	Size          int
// 	NonNull       bool
// }
//
// type TableConfig struct {
// 	Name        string
// 	Columns     []*ColumnConfig
// 	ForeignKeys []*ForeignKeys
// 	Indexes     []*Index
// }
