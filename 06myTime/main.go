package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to story of time in go lang")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Printf("presentTime type is %T\n", presentTime)
	formattedPresentTime := presentTime.Format("02-01-2006 Monday  15:04:05")
	fmt.Println(formattedPresentTime)
	fmt.Printf("formattedPresentTime type is %T\n", formattedPresentTime)

	createdTime := time.Date(2024, time.October, 19, 2, 45, 20, 0, time.UTC)
	fmt.Printf("createdTime type is %T", createdTime)
	fmt.Println(createdTime)
}
