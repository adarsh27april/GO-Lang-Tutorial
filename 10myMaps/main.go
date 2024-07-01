package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps = HastTable = keyValueTable in other langs")

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["java"] = "Java"
	languages["rb"] = "Ruby"

	fmt.Println(languages)

	fmt.Println(languages["py"])
	delete(languages, "py")
	fmt.Println(languages)

	// loops in go lang
	for _, value := range languages {
		fmt.Printf("val == '%v'\n", value)
	}
}
