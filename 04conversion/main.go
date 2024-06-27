package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Println("Please rate our pizza b/w 1 & 5")

	reader := bufio.NewReader(os.Stdin)

	ip, _ := reader.ReadString('\n')
	fmt.Println("Thanks for your rating : ", ip)
	// converting rating to number

	numRating, err := strconv.ParseFloat(strings.TrimSpace(ip), 64)
	if err != nil {
		fmt.Println("some error ->", err)
	} else {
		fmt.Printf("var type is %T", numRating)
	}

}
