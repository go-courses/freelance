package db

import (
	"database/sql"
	"fmt"

	// This line is must for working PostgreSQL database

	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	_ "github.com/mattes/migrate/source/file"

	"github.com/jmoiron/sqlx"

	"github.com/go-courses/freelance/config"
)

// PgSQL provides api for work with mysql database
type PgSQL struct {
	conn        *sqlx.DB
	connmigrate *sql.DB
}

// NewPgSQL creates a new instance of database API
func NewPgSQL(c *config.FreelanceConfig) (*PgSQL, error) {
	connmigrate, err := sql.Open("postgres", c.DatabaseURL)
	if err != nil {
		return nil, err
	}

	conn := sqlx.NewDb(connmigrate, "postgres")

	m := &PgSQL{}
	m.conn = conn
	m.connmigrate = connmigrate
	return m, nil
}

// MigrateUp - create tables
func (m *PgSQL) MigrateUp() error {
	driver, _ := postgres.WithInstance(m.connmigrate, &postgres.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		"file:///home/meder/go/src/github.com/go-courses/freelance/migrations/pgsql",
		"postgres",
		driver,
	)

	if err != nil {
		fmt.Println(err, "file not foundd")
	}

	err = migration.Up()
	if err != nil {
		return err
	}
	return nil
}

// MigrateDown - delete tables
func (m *PgSQL) MigrateDown() error {
	driver, _ := postgres.WithInstance(m.connmigrate, &postgres.Config{})
	migration, _ := migrate.NewWithDatabaseInstance(
		"file:///home/meder/go/src/github.com/go-courses/freelance/migrations/pgsql",
		"postgres",
		driver,
	)

	err := migration.Down()
	if err != nil {
		return err
	}
	return nil
}
