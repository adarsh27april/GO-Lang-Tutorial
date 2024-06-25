package main

import "fmt"

func main() {
	fmt.Println("Welcome to story of pointers in Go Lang")

	var ptr *int // ptr is a pointer which store address of int data type
	fmt.Println("value of ptr is  : ", ptr)

	num := 24
	var ptr1 = &num
	fmt.Println("value of ptr1 is  : ", ptr1)
	fmt.Println("value of ptr1 is  : ", *ptr1)
	*ptr1 = *ptr1 * 3
	fmt.Println("new value of num is  : ", num)
}
