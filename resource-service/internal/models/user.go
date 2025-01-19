package models

type User struct {
	Id       int
	Username string
	Email    string
	Roles    []string
}
