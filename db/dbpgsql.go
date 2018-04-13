package db

import (
	"time"

	"github.com/go-courses/freelance/model"
	"github.com/lib/pq"
	"github.com/mattes/migrate/source/file"

	"github.com/jmoiron/sqlx"
)

// PgSQL provides api for work with PgSQL database
type PgSQL struct {
	conn *sqlx.DB
}

type Config struct {
	ConnectString string
}

// NewPgSQL creates a new instance of database API
func NewPgSQL(cfg Config) (*PgSQL, error) {
	if conn, err := sqlx.Connect("postgres", cfg.ConnectString); err != nil {
		return nil, err
	} else {
		p := &PgSQL{conn: conn}
		if err := p.conn.Ping(); err != nil {
			return nil, err
		}
		return p, nil
	}
}

// CreateUser creates user entry in database
func (m *PgSQL) CreateUser(s model.User) (model.User, error) {
	res, err := m.conn.Exec(
		"INSERT INTO `users` (name, utype, balance) VALUES (?, ?, ?)",
		s.Name, s.UserType, s.Balance
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
func (m *PgSQL) SelectUser(id int64) (model.User, error) {
	var s model.User
	err := m.conn.Get(&s, "SELECT * FROM `user` WHERE id=?", id)
	return s, err
}

// ListUsers returns array of users entries from database
func (m *PgSQL) ListUsers() ([]model.User, error) {
	users := []model.User{}
	err := m.conn.Select(&users, "SELECT * FROM `users`")
	return users, err
}

// UpdateUser updates user entry in database
func (m *PgSQL) UpdateUser(s model.User) (model.User, error) {
	_, err := m.conn.Exec(
		"UPDATE `users` SET name=?, utype=?, balance=? WHERE id=?",
		s.Name, s.UserType, s.Balance, s.ID,
	)
	if err != nil {
		return s, err
	}
	var i model.User
	err = m.conn.Get(&i, "SELECT * FROM `users` WHERE id=?", s.ID)
	return i, err
}

// DeleteUser deletes user entry from database
func (m *PgSQL) DeleteUser(id int64) error {
	_, err := m.conn.Exec("DELETE FROM `users` WHERE id=?", id)
	return err
}


// CreateTask creates task entry in database
func (m *PgSQL) CreateTask(s model.Task) (model.Task, error) {
	res, err := m.conn.Exec(
		"INSERT INTO `tasks` (description, price, status) VALUES (?, ?, ?)",
		s.Description, s.Price, s.Status
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
func (m *PgSQL) SelectTask(id int64) (model.Task, error) {
	var s model.Task
	err := m.conn.Get(&s, "SELECT * FROM `tasks` WHERE id=?", id)
	return s, err
}

// ListTasks returns array of tasks entries from database
func (m *PgSQL) ListTasks() ([]model.Task, error) {
	tasks := []model.Task{}
	err := m.conn.Select(&tasks, "SELECT * FROM `tasks`")
	return tasks, err
}

// UpdateTask updates task entry in database
func (m *PgSQL) UpdateTask(s model.Task) (model.Task, error) {
	_, err := m.conn.Exec(
		"UPDATE `tasks` SET description=?, price=?, status=? WHERE id=?",
		s.Description, s.Price, s.Status, s.ID,
	)
	if err != nil {
		return s, err
	}
	var i model.Task
	err = m.conn.Get(&i, "SELECT * FROM `tasks` WHERE id=?", s.ID)
	return i, err
}

// DeleteTask deletes task entry from database
func (m *PgSQL) DeleteTask(id int64) error {
	_, err := m.conn.Exec("DELETE FROM `tasks` WHERE id=?", id)
	return err
}



// CreateBilling creates billing entry in database
func (m *PgSQL) CreateBilling(s model.Billing) (model.Billing, error) {
	res, err := m.conn.Exec(
		"INSERT INTO `billings` (sender, reciever, amount, time_bill, task_id, btype) VALUES (?, ?, ?, ?, ?,?)",
		s.Sender, s.Reciever, s.amount, s.time.Now().UTC(), s.TaskID, s.BillingType
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
func (m *PgSQL) SelectBilling(id int64) (model.Billing, error) {
	var s model.Billing
	err := m.conn.Get(&s, "SELECT * FROM `billings` WHERE id=?", id)
	return s, err
}

// ListBillings returns array of Billing entries from database
func (m *PgSQL) ListBillings() ([]model.Billing, error) {
	billings := []model.Billing{}
	err := m.conn.Select(&billings, "SELECT * FROM `billings`")
	return billings, err
}

// UpdateBilling updates billing entry in database
func (m *PgSQL) UpdateBilling(s model.Billing) (model.Billing, error) {
	_, err := m.conn.Exec(
		"UPDATE `billings` SET sender=?, reciever=?, amount=?, time_bill=?, task_id=?, btype=? WHERE id=?",
		s.Sender, s.Reciever, s.amount, s.TimeBill, s.TaskID, s.BillingType, s.ID,
	)
	if err != nil {
		return s, err
	}
	var i model.Billing
	err = m.conn.Get(&i, "SELECT * FROM `billings` WHERE id=?", s.ID)
	return i, err
}

// DeleteBilling deletes billing entry from database
func (m *PgSQL) DeleteBilling(id int64) error {
	_, err := m.conn.Exec("DELETE FROM `billings` WHERE id=?", id)
	return err
}