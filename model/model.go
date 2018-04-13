package model

type Model struct {
	db
}

func New(db db) *Model {
	return &Model{
		db: db,
	}
}

func (m *Model) Users() ([]*User, error) {
	return m.SelectUsers()
}

func (m *Model) Tasks() ([]*Task, error) {
	return m.SelectTasks()
}

func (m *Model) Billings() ([]*Billing, error) {
	return m.SelectBillings()
}
