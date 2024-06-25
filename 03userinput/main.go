package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	// do note that here `NewReader` has `N` & `Stdin` has `S` hence they are public
	fmt.Println("Enter rating for our pizza :") // println auto adds a \n at end

	// `, ok syntax` || `err ok syntax`

	/*
		The paradigm of the language is that it expects err to be treated like a var/bool
	*/

	input, _ := reader.ReadString('\n') // reads until the first occurrence of delimeter in the input.
	fmt.Println("Thanks for rating", input)
	fmt.Printf("Type of rating is %T", input)
	// note that err is not used here so we can use `_` as a var name here and it will not show `Declared but not user` error
}
