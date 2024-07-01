package main

import "fmt"

func main() {
	fmt.Println("Structs in GO Lang")
	adarsh := User{"Adarsh", "adarsh@go.dev", true, 24}
	adarsh.GetStatus()
	adarsh.UpdateEmail("test@go.dev")

	fmt.Println("original adarsh.Email is : ", adarsh.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("user  : ", u.Status)
}

func (u User) UpdateEmail(email string) {
	u.Email = email
	fmt.Println("Updated Email is : ", u.Email)
}
