package model

type db interface {
	SelectUsers() ([]*User, error)
	SelectTasks() ([]*Task, error)
	SelectBillings() ([]*Billing, error)
}
