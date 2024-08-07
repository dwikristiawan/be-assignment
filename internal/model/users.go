package model

type User struct {
	ID       int    `json:"id" db:"id"`
	Password string `json:"password" db:"password"`
}
