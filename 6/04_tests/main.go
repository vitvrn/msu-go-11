package main

func main() {

}

func Sum(a int, b int) int {
	return a + b
}

func Devision(a int, b int) int {
	return a / b
}

type User struct {
	Name string
	Age int
}

func AddUser(users *[]User, name string, age int) {
	user := User{
		Name: name,
		Age: age,
	}

	*users = append(*users, user)
}