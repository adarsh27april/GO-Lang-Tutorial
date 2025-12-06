package main

import (
	"fmt"
)

func main() {
	// fmt.Println("Starting Textio Server")
	// var user string = "adarsh"
	// println("user is ", user)

	// test("Lane,", " happy birthday!")
	// test("Elon,", " hope that Tesla thing works out")
	// test("Go", " is fantastic")

	// fmt.Println(time.Now())
	fizzBuzz(15)
}

func concat(s1, s2 string) string {
	return s1 + s2
}

// don't touch below this line

func test(s1 string, s2 string) {
	fmt.Println(concat(s1, s2))
}

func fizzBuzz(n int32) {
	// Write your code here
	var i int32
	for i = 1; i <= n; i++ {
		s := ""
		if i%3 == 0 {
			s += "Fizz"
		} else if i%5 == 0 {
			s += "Buzz"
		} else {
			s = fmt.Sprintf("%d", i)
		}
		fmt.Println(s)
	}
}
