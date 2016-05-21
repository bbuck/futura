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
	"fmt"
	"os"

	"github.com/bbuck/futura/log"
	"github.com/jackc/pgx"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func init() {
	Map["pgx"] = PgxAdapter{}
}

// migration constants
const (
	PgxUpComment   = "-- +MIGRATE UP"
	PgxDownComment = "-- +MIGRATE DOWN"
)

// PgxAdapter uses Jack C's pgx Postgres Adapter to power migrations.
type PgxAdapter struct{}

// CreateDatabase will create the given database for the the current environemnt
// unless it already exists.
func (p PgxAdapter) CreateDatabase(name string) error {
	conn := pgxConnection(false)
	defer conn.Close()
	_, err := conn.Exec(fmt.Sprintf("CREATE DATABASE %s", name))

	return err
}

// CreateVersionTable will create a version table (or noop) if required by the
// Adapter.
func (p PgxAdapter) CreateVersionTable() error {
	return nil
}

// RunRawMigration will execute a query in the context of the database for the
// current environment.
func (p PgxAdapter) RunRawMigration(query string) error {
	return nil
}

// RawFileExtension returns "sql" as the file extension for Raw queries using
// this adapter.
func (p PgxAdapter) RawFileExtension() string {
	return "sql"
}

// RawFileParser returns a FileParser object that will parse raw migration files
// for the given Adapter.
func (p PgxAdapter) RawFileParser(filename string) (*FileParser, error) {
	return NewFileParser(filename, p.UpComment(), p.DownComment())
}

// DefaultConfigString returns the basic (or extensive) settings to fill out
// an automatically generated Futurafile.
func (PgxAdapter) DefaultConfigString() string {
	return `adapter = "pgx"

migration_path = "db/migrate"

database = "your_project_dev"
# hostname = "localhost"
# port = 5432
# user = "username"
# password = "password"`
}

// Description returns a short descriptions for listing out the supported
// adapters.
func (PgxAdapter) Description() string {
	return "The pgx adapter uses jackc/pgx to connect to Postgres."
}

// CommentStart returns the -- that initiate a comment in postgres
func (PgxAdapter) CommentStart() string {
	return "--"
}

// UpComment returns the comment marking an up migration
func (PgxAdapter) UpComment() string {
	return "-- +MIGRATE UP"
}

// DownComment returns the comment marking a down migration
func (PgxAdapter) DownComment() string {
	return "-- +MIGRATE DOWN"
}

// connect to postgres
func pgxConnection(withDb bool) *pgx.Conn {
	cfg := pgx.ConnConfig{Database: "postgres"}
	settings := viper.GetStringMap(viper.GetString("env"))
	if val, ok := settings["hostname"]; ok {
		cfg.Host = cast.ToString(val)
	}
	if val, ok := settings["port"]; ok {
		cfg.Port = uint16(cast.ToInt(val))
	}
	if val, ok := settings["database"]; ok && withDb {
		cfg.Database = cast.ToString(val)
	}
	if val, ok := settings["user"]; ok {
		cfg.User = cast.ToString(val)
	}
	if val, ok := settings["password"]; ok {
		cfg.Password = cast.ToString(val)
	}

	conn, err := pgx.Connect(cfg)
	if err != nil {
		log.Errorf("Failed to establish Postgres connection: %s\n", err.Error())
		os.Exit(4)
	}

	return conn
}
