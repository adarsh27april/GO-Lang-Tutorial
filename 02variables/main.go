package main

import "fmt"

const LoginToken = "hiurhgfuiehrnvi"

func main() {
	var username string = "adarsh"
	fmt.Println(username)
	fmt.Printf("var is of type : %T \n", username)

	var smallFloat float64 = 256.1234567890
	fmt.Println(smallFloat)
	fmt.Printf("var is of type : %T \n", smallFloat)

	fmt.Println(LoginToken)
	fmt.Printf("var is of type : %T", LoginToken)
}
