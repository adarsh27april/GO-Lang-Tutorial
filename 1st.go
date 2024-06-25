package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println("Starting Textio Server")
	// var user string = "adarsh"
	// println("user is ", user)

	// test("Lane,", " happy birthday!")
	// test("Elon,", " hope that Tesla thing works out")
	// test("Go", " is fantastic")

	fmt.Println(time.Now())
}

func concat(s1, s2 string) string {
	return s1 + s2
}

// don't touch below this line

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}
