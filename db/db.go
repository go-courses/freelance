package db

import "github.com/go-courses/freelance/model"

// DB ...
type DB interface {
	CreateUser(s model.User) (model.User, error)
	SelectUser(id int64) (model.User, error)
	ListUsers() ([]model.User, error)
	UpdateUser(s model.User) (model.User, error)
	DeleteUser(id int64) error

	CreateTask(s model.Task) (model.Task, error)
	SelectTask(id int64) (model.Task, error)
	ListTasks() ([]model.Task, error)
	UpdateTask(s model.Task) (model.Task, error)
	DeleteTask(id int64) error

	CreateBilling(s model.Billing) (model.Billing, error)
	SelectBilling(id int64) (model.Billing, error)
	ListBillings() ([]model.Billing, error)
	UpdateBilling(s model.Billing) (model.Billing, error)
	DeleteBilling(id int64) error

	MigrateUp(string) error
	MigrateDown(string) error
}
