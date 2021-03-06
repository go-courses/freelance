package db

import (
	"database/sql"

	"github.com/go-courses/freelance/config"
	"github.com/go-courses/freelance/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mattes/migrate"
	"github.com/mattes/migrate/database/postgres"
	"github.com/pkg/errors"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "dbuser_f"
	password = "dbpass_f"
	dbname   = "freelance"
)

// PgSQL provides api for work with PgSQL database
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

// CreateUser creates user entry in database
func (m *PgSQL) CreateUser(s model.User) (model.User, error) {
	sqlStatement := `INSERT INTO users (name, utype, balance)
	VALUES ($1, $2, $3)
	RETURNING id`
	id := 0
	err := m.conn.QueryRow(sqlStatement, s.Name, s.UserType, s.Balance).Scan(&id)
	if err != nil {
		panic(err)
	}
	s.ID = int64(id)
	return s, nil
}

// SelectUser selects user entry from database
func (m *PgSQL) SelectUser(id int64) (model.User, error) {
	var s model.User
	sqlStatement := `SELECT * FROM users WHERE id=$1;`
	err := m.conn.Get(&s, sqlStatement, id)
	return s, err
}

// ListUsers returns array of users entries from database
func (m *PgSQL) ListUsers() ([]model.User, error) {
	users := []model.User{}
	err := m.conn.Select(&users, `SELECT * FROM users`)
	return users, err
}

// UpdateUser updates user entry in database
func (m *PgSQL) UpdateUser(s model.User) (model.User, error) {
	sqlStatement := `UPDATE users
	SET name = $2, utype = $3, balance = $4
	WHERE id = $1;`
	tx := m.conn.MustBegin()
	tx.MustExec(sqlStatement, s.ID, s.Name, s.UserType, s.Balance)
	err := tx.Commit()

	if err != nil {
		return s, err
	}
	var i model.User
	err = m.conn.Get(&i, `SELECT * FROM users WHERE id=$1`, s.ID)
	return i, err
}

// DeleteUser deletes user entry from database
func (m *PgSQL) DeleteUser(id int64) error {
	sqlStatement := `DELETE FROM users WHERE id = $1;`
	_, err := m.conn.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	return err
}

// CreateTask creates task entry in database
func (m *PgSQL) CreateTask(s model.Task) (model.Task, error) {
	sqlStatement := `INSERT INTO users (description, price, status)
	VALUES ($1, $2, $3)
	RETURNING id`
	id := 0
	err := m.conn.QueryRow(sqlStatement, s.Description, s.Price, s.Status).Scan(&id)
	if err != nil {
		panic(err)
	}
	s.ID = int64(id)
	return s, nil
}

// SelectTask selects task entry from database
func (m *PgSQL) SelectTask(id int64) (model.Task, error) {
	var s model.Task
	sqlStatement := `SELECT * FROM tasks WHERE id=$1;`
	err := m.conn.Get(&s, sqlStatement, id)
	return s, err
}

// ListTasks returns array of tasks entries from database
func (m *PgSQL) ListTasks() ([]model.Task, error) {
	tasks := []model.Task{}
	err := m.conn.Select(&tasks, `SELECT * FROM tasks`)
	return tasks, err
}

// UpdateTask updates task entry in database
func (m *PgSQL) UpdateTask(s model.Task) (model.Task, error) {
	sqlStatement := `UPDATE tasks
	SET description = $2, price = $3, status = $4
	WHERE id = $1;`
	_, err := m.conn.Exec(sqlStatement, s.ID, s.Description, s.Price, s.Status)

	if err != nil {
		return s, err
	}
	var i model.Task
	err = m.conn.Get(&i, `SELECT * FROM tasks WHERE id=$1`, s.ID)
	return i, err
}

// DeleteTask deletes task entry from database
func (m *PgSQL) DeleteTask(id int64) error {
	sqlStatement := `DELETE FROM tasks
	WHERE id = $1;`
	_, err := m.conn.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	return err
}

// CreateBilling creates billing entry in database
func (m *PgSQL) CreateBilling(s model.Billing) (model.Billing, error) {
	sqlStatement := `INSERT INTO billings (sender, reciever, amount, time_bill, task_id, btype)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id`
	id := 0
	err := m.conn.QueryRow(sqlStatement, s.Sender, s.Reciever, s.Amount, s.TimeBill, s.TaskID, s.BillingType).Scan(&id)
	if err != nil {
		panic(err)
	}
	s.ID = int64(id)
	return s, nil

}

// SelectBilling selects billing entry from database
func (m *PgSQL) SelectBilling(id int64) (model.Billing, error) {
	var s model.Billing
	sqlStatement := `SELECT * FROM billings WHERE id=$1;`
	err := m.conn.Get(&s, sqlStatement, id)
	return s, err

}

// ListBillings returns array of Billing entries from database
func (m *PgSQL) ListBillings() ([]model.Billing, error) {
	billings := []model.Billing{}
	err := m.conn.Select(&billings, `SELECT * FROM billings`)
	return billings, err
}

// UpdateBilling updates billing entry in database
func (m *PgSQL) UpdateBilling(s model.Billing) (model.Billing, error) {
	sqlStatement := `UPDATE billings
	SET sender=$1, reciever=$2, amount=$3, time_bill=$4, task_id=$5, btype=$6
	WHERE id = $7;`
	tx := m.conn.MustBegin()
	tx.MustExec(sqlStatement, s.Sender, s.Reciever, s.Amount, s.TimeBill, s.TaskID, s.BillingType, s.ID)
	err := tx.Commit()

	if err != nil {
		return s, err
	}
	var i model.Billing
	err = m.conn.Get(&i, `SELECT * FROM billings WHERE id=$1`, s.ID)
	return i, err
}

// DeleteBilling deletes billing entry from database
func (m *PgSQL) DeleteBilling(id int64) error {
	sqlStatement := `DELETE FROM billings
	WHERE id = $1;`
	_, err := m.conn.Exec(sqlStatement, id)
	if err != nil {
		panic(err)
	}
	return err
}

// MigrateUp - create tables
func (m *PgSQL) MigrateUp(migrateFolder string) error {
	driver, _ := postgres.WithInstance(m.connmigrate, &postgres.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		migrateFolder,
		"postgres",
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
func (m *PgSQL) MigrateDown(migrateFolder string) error {
	driver, _ := postgres.WithInstance(m.connmigrate, &postgres.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		migrateFolder,
		"postgres",
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
