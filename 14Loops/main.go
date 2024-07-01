package main

import (
	"fmt"
	"time"
)

func main() {
	days := []string{time.Monday.String(), time.Tuesday.String(), time.Wednesday.String(), time.Thursday.String(), time.Friday.String(), time.Saturday.String()}
	fmt.Println(days)
	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	// i returns index and not the actual slice element
	// 	fmt.Println(days[i])
	// }

	// for _, day := range days {
	// 	// to not use i use `_`
	// 	fmt.Printf(" day is %v\n", day)
	// }

	rougueValue := 1
	for rougueValue < 10 {
		fmt.Println("rougueVal : ", rougueValue)
		rougueValue++
	}
}
