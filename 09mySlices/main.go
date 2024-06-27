package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices in Go Lang")
	// inserting into slice
	// var fruitList = []string{"Mango", "Apple", "Bananas"}
	// fmt.Printf("type of fruitList: %T\n", fruitList)
	// fmt.Println(fruitList)
	// fruitList = append(fruitList, "Peach", "Kiwi")
	// fmt.Println("updated: ", fruitList)
	// fruitList = append(fruitList[1:3])
	// fmt.Println(fruitList) // [Apple Bananas]
	// it's like
	// [startIndex : endIndex)
	// slices using make
	// highScores := make([]int, 4)
	// highScores[0] = 999
	// highScores[1] = 888
	// highScores[2] = 777
	// highScores[3] = 666
	// fmt.Println(highScores)
	// highScores[4] = 555
	// but if I use append here
	// highScores = append(highScores, 555, 444, 333, 222)
	// fmt.Println(highScores)
	// sort.Ints(highScores)
	// fmt.Println(highScores)

	// remove a value from slices based on index
	var technologies = []string{"reactjs", "nestjs", "gokit", "django", "flask"}
	fmt.Println(technologies)
	indexToRemove := 3 // i.e., remove `django`
	technologies = append(technologies[:indexToRemove], technologies[indexToRemove+1:]...)
	fmt.Println(technologies)
}
