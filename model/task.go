package model

type Task struct {
	ID          int64  `db:"id"`
	Description string `db:"description"`
	Price       int64  `db:"price"`
	Status      string `db:"status"`
}
