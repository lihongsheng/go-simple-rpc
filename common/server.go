package common

type User struct {
	Id int
	Name string
}

type Server interface {
	GetUser(id int) User
}
