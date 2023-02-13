package importTest

import "fmt"

func (DB Database) Registration(u User) *Database {
	DB.Users = append(DB.Users, u)
	fmt.Println("Successfully registered")
	return &DB
}
