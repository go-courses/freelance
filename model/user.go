package model

type User struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Type    string `db:"utype"`
	Balanse int64  `db:"balance"`
}
