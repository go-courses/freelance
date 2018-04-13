package model

// Task ....
type Task struct {
	ID          int64  `db:"id"`
	Description string `db:"description"`
	Price       int32  `db:"price"`
	Status      string `db:"status"`
}
