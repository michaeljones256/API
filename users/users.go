package users

import "fmt"

//type User struct {
// 	name string
// 	age  int
// }

// func NewUser(name string, age int) *User {
// 	return &User{name: name, age: age}
// }

type UserService struct {
	store UserStore
}

func NewUserService(s UserStore) *UserService {
	return &UserService{store: s}
}

type UserStore interface {
	Insert(item interface{})
	Get(id int)
}

func (u *UserService) CreateUser() {
	fmt.Println("Create a user")
	u.store.Get()

}
