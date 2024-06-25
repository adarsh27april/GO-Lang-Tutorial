package main

import "fmt"

func main() {
	fmt.Println("Story of arrays in go lang")

	var fruitList [5]string
	fruitList[0] = "Mango"
	fruitList[1] = "Apple"
	fruitList[2] = "Bananas"
	fruitList[4] = "Peach"
	fmt.Println("fruitList is: ", fruitList)
	fmt.Println("fruitList length is: ", len(fruitList))

	var vegList = [3]string{"potato", "mushroom", "beans"}
	fmt.Println(vegList)
}
