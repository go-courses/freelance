package db

import (
	"database/sql"
	"go/build"
	"os"

	// This line is must for working MySQL database

	_ "github.com/go-sql-driver/mysql"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/mysql"
	_ "github.com/mattes/migrate/source/file"
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"

	"github.com/go-courses/freelance/config"
)

// MySQL provides api for work with mysql database
type MySQL struct {
	conn        *sqlx.DB
	connmigrate *sql.DB
}

func getPath() string {
	gp := os.Getenv("GOPATH")
	if gp == "" {
		gp = build.Default.GOPATH
	}
	return gp
}

// NewMySQL creates a new instance of database API
func NewMySQL(c *config.FreelanceConfig) (*MySQL, error) {
	connmigrate, err := sql.Open("mysql", c.DatabaseURL)
	if err != nil {
		return nil, err
	}

	conn := sqlx.NewDb(connmigrate, "mysql")

	m := &MySQL{}
	m.conn = conn
	m.connmigrate = connmigrate
	return m, nil
}

// MigrateUp - create tables
func (m *MySQL) MigrateUp() error {
	driver, _ := mysql.WithInstance(m.connmigrate, &mysql.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		"file://"+getPath()+"/src/github.com/go-courses/freelance/migrations/mysql",
		"mysql",
		driver,
	)

	if err != nil {
		return errors.Wrap(err, "migration file not found")
	}

	err = migration.Up()
	if err != nil {
		return errors.Wrap(err, "migration Up error")
	}
	return nil
}

// MigrateDown - delete tables
func (m *MySQL) MigrateDown() error {
	driver, _ := mysql.WithInstance(m.connmigrate, &mysql.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		"file://"+getPath()+"/src/github.com/go-courses/freelance/migrations/mysql",
		"mysql",
		driver,
	)

	if err != nil {
		return errors.Wrap(err, "migration file not found")
	}

	err = migration.Down()
	if err != nil {
		return errors.Wrap(err, "migration Down error")
	}
	return nil
}
