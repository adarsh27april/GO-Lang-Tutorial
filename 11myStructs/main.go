package main

import "fmt"

func main() {
	fmt.Println("Structs in GO Lang")

	adarsh := User{"Adarsh", "adarsh@go.dev", true, 24}

	fmt.Println(adarsh)
	fmt.Printf("inside quotes %+v\n", adarsh)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
