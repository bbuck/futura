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

	"github.com/jackc/pgx"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
)

func init() {
	Map["pgx"] = PgxAdapter{}
}

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
	return NewFileParser(filename, "-- +MIGRATE UP", "-- +MIGRATE DOWN")
}

// DefaultConfigString returns the basic (or extensive) settings to fill out
// an automatically generated Futurafile.
func (p PgxAdapter) DefaultConfigString() string {
	return `adapter = "pgx"

database = "your_project_dev"
# hostname = "localhost"
# port =
# user = "username"
# password = "password"`
}

// Description returns a short descriptions for listing out the supported
// adapters.
func (p PgxAdapter) Description() string {
	return "The pgx adapter uses jackc/pgx to provide an interface with the Postgres database server."
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
		fmt.Fprintf(os.Stderr, "Failed to establish Postgres connection: %s\n", err.Error())
		os.Exit(4)
	}

	return conn
}
