package main

// import(
//
// )

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func newUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

type Database struct {
	Users []*User
}

func newDatabase() *Database {
	return &Database{}
}

func (db *Database) AddUser(user *User) {
	db.Users = append(db.Users, user)
}
