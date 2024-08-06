package main

import (
	"fmt"
	"io"
	"net/http"
)

// const url = "https://www.google.com/robots.txt"
// const url = "https://lco.dev"
const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Handling web requests")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close() // it is caller's responsibility to close the connection
	// adding defer to this statement ensures that it will be executed at the end of the function, i.e., after everything else has been executed

	fmt.Printf("Response is of Type : %T\n", res) // Response is of Type : *http.Response
	// note that it is returning a pointer to the acutal Response and not a copy of it, so it can be manipulated further.

	databytes, err := io.ReadAll(res.Body) // ReadAll reads from r until an error or EOF and returns the data it read

	if err != nil {
		panic(err)
	}
	// fmt.Println("databytes : ", databytes)

	content := string(databytes) // it will give us the entire content, text/html whatever is there
	fmt.Println("content : ", content)
}
