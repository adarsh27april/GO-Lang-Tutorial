package main

import "fmt"

func main() {
	fmt.Println("functions")
	fmt.Println(proAdder(1, 2, 3, 4, 5))

	fmt.Println(twoReturner(2))
}

// PRO Function
func proAdder(values ...int) int {
	// values will be accepted as a slice
	total := 0
	for _, value := range values {
		total += value
	}
	return total
}

// returning 2 values from function
func twoReturner(a int) (int, string) {
	return a, "funny func"
}
