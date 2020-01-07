package model

// User .
type User struct {
	ID   int    `db:"Id"`
	Name string `db:"Name"`
	Timestamp
}
