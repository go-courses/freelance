package model

// Task ....
type Task struct {
	ID          int64  `db:"id"`
	Description string `db:"description"`
	Creator     int64  `db:"creator"`
	Executor    int64  `db:"executor"`
	Price       *Money  `db:"price"`
	Status      string `db:"status"`
}
