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
	"github.com/go-courses/freelance/model"
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

// CreateUser creates user entry in database
func (m *MySQL) CreateUser(s model.User) (model.User, error) {
	res, err := m.conn.Exec(
		"INSERT INTO `users` (`name`, `utype`, `balance`) VALUES (?, ?, ?)", s.Name, s.UserType, s.Balance,
	)
	if err != nil {
		return s, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return s, err
	}
	s.ID = id
	return s, nil
}

// SelectUser selects user entry from database
func (m *MySQL) SelectUser(id int64) (model.User, error) {
	var s model.User
	err := m.conn.Get(&s, "SELECT * FROM `users` WHERE id=?", id)
	return s, err
}

// ListUsers returns array of users entries from database
func (m *MySQL) ListUsers() ([]model.User, error) {
	users := []model.User{}
	err := m.conn.Select(&users, "SELECT * FROM `users`")
	return users, err
}

// UpdateUser updates user entry in database
func (m *MySQL) UpdateUser(s model.User) (model.User, error) {
	tx := m.conn.MustBegin()
	tx.MustExec(
		"UPDATE `users` SET `name` = ?, `utype` = ?, `balance` = ? WHERE `id` = ?",
		s.Name, s.UserType, s.Balance, s.ID,
	)
	err := tx.Commit()

	if err != nil {
		return s, err
	}
	var i model.User
	err = m.conn.Get(&i, "SELECT * FROM `users` WHERE id=?", s.ID)
	return i, err
}

// DeleteUser deletes user entry from database
func (m *MySQL) DeleteUser(id int64) error {
	_, err := m.conn.Exec("DELETE FROM `users` WHERE id=?", id)
	return err
}

// CreateTask creates task entry in database
func (m *MySQL) CreateTask(s model.Task) (model.Task, error) {
	res, err := m.conn.Exec(
		"INSERT INTO `tasks` (description, creator, executor, price, status) VALUES (?, ?, ?, ?, ?)",
		s.Description, s.Creator, s.Executor, s.Price, s.Status,
	)
	if err != nil {
		return s, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return s, err
	}
	s.ID = id
	return s, nil
}

// SelectTask selects task entry from database
func (m *MySQL) SelectTask(id int64) (model.Task, error) {
	var s model.Task
	err := m.conn.Get(&s, "SELECT * FROM `tasks` WHERE id=?", id)
	return s, err
}

// ListTasks returns array of tasks entries from database
func (m *MySQL) ListTasks() ([]model.Task, error) {
	tasks := []model.Task{}
	err := m.conn.Select(&tasks, "SELECT * FROM `tasks`")
	return tasks, err
}

// UpdateTask updates task entry in database
func (m *MySQL) UpdateTask(s model.Task) (model.Task, error) {
	_, err := m.conn.Exec(
		"UPDATE `tasks` SET description=?, creator=?, executor=?, price=?, status=? WHERE id=?",
		s.Description, s.Creator, s.Executor, s.Price, s.Status, s.ID,
	)
	if err != nil {
		return s, err
	}
	var i model.Task
	err = m.conn.Get(&i, "SELECT * FROM `tasks` WHERE id=?", s.ID)
	return i, err
}

// DeleteTask deletes task entry from database
func (m *MySQL) DeleteTask(id int64) error {
	_, err := m.conn.Exec("DELETE FROM `tasks` WHERE id=?", id)
	return err
}

// CreateBilling creates billing entry in database
func (m *MySQL) CreateBilling(s model.Billing) (model.Billing, error) {
	res, err := m.conn.Exec(
		"INSERT INTO `billings` (sender, reciever, amount, time_bill, task_id, btype) VALUES (?, ?, ?, ?, ?,?)",
		s.Sender, s.Reciever, s.Amount, s.TimeBill, s.TaskID, s.BillingType,
	)
	if err != nil {
		return s, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return s, err
	}
	s.ID = id
	return s, nil
}

// SelectBilling selects billing entry from database
func (m *MySQL) SelectBilling(id int64) (model.Billing, error) {
	var s model.Billing
	err := m.conn.Get(&s, "SELECT * FROM `billings` WHERE id=?", id)
	return s, err
}

// ListBillings returns array of Billing entries from database
func (m *MySQL) ListBillings() ([]model.Billing, error) {
	billings := []model.Billing{}
	err := m.conn.Select(&billings, "SELECT * FROM `billings`")
	return billings, err
}

// UpdateBilling updates billing entry in database
func (m *MySQL) UpdateBilling(s model.Billing) (model.Billing, error) {
	tx := m.conn.MustBegin()
	tx.MustExec(
		"UPDATE `billings` SET `sender` = ?, `reciever` = ?, `amount` = ?, `time_bill` = ?, `task_id` = ?, `btype` = ? WHERE `id` = ?",
		s.Sender, s.Reciever, s.Amount, s.TimeBill, s.TaskID, s.BillingType, s.ID,
	)
	err := tx.Commit()
	if err != nil {
		return s, errors.Wrap(err, "not updating")
	}
	var i model.Billing
	err = m.conn.Get(&i, "SELECT * FROM `billings` WHERE id=?", s.ID)
	return i, errors.Wrapf(err, "Not selected %d -- %s", s.ID, s.BillingType)
}

// DeleteBilling deletes billing entry from database
func (m *MySQL) DeleteBilling(id int64) error {
	_, err := m.conn.Exec("DELETE FROM `billings` WHERE id=?", id)
	return err
}

// MigrateUp - create tables
func (m *MySQL) MigrateUp(migrateFolder string) error {
	driver, _ := mysql.WithInstance(m.connmigrate, &mysql.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		migrateFolder,
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
func (m *MySQL) MigrateDown(migrateFolder string) error {
	driver, _ := mysql.WithInstance(m.connmigrate, &mysql.Config{})
	migration, err := migrate.NewWithDatabaseInstance(
		migrateFolder,
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
