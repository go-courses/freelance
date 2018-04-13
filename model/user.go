package model

type User struct {
	ID       int64  `db:"id"`
	Name     string `db:"name"`
	UserType string `db:"utype"`
	Balance  int    `db:"balance"`
}
