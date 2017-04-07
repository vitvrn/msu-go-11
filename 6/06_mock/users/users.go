package users

type User struct {
	ID int
	Name string
	Age int
}

type Users interface {
	Get(id int) User
	Put(id int, name string, age int)
}

func DoSomethingWithUsers(users *Users) {
	(*users).Put(500, "name", 7)
}

