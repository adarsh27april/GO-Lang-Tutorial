package main

import (
	"fmt"
	"net/url"
)

// const url1 string = "https://jsonplaceholder.typicode.com/todos/1"
const url1 string = "https://jsonplaceholder.typicode.com/comments?postId=1&userId=234"

func main() {
	fmt.Println("handling urls in go lang")

	// parsing the url
	result, _ := url.Parse(url1)

	fmt.Println("scheme : ", result.Scheme)         // https
	fmt.Println("host : ", result.Host)             // jsonplaceholder.typicode.com
	fmt.Println("path : ", result.Path)             // /comments
	fmt.Println("port : ", result.Port())           //
	fmt.Println("query params : ", result.RawQuery) // postId=1

	// parsing the query params
	queryParams := result.Query()
	fmt.Println("queryParams : ", queryParams)             // queryParams :  map[postId:[1]]
	fmt.Printf("Type of queryParams is %T\n", queryParams) // url.Values
	// it returns a map
	fmt.Println(`queryParams["postId"] :`, queryParams["postId"])

	for _, val := range queryParams { // the params in teaversing url.Values is key, value basically an interface
		fmt.Println("param is : ", val)
	}

	partsOfUrl := &url.URL{
		// & denotes that the acutal pointer to the url.URL is being passed so it can acutally be manipulated
		// URL is a struct
		Scheme:   "https",
		Host:     "jsonplaceholder.typicode.com",
		Path:     "/comments",
		RawQuery: "postId=1&userId=234",
	}
	fmt.Printf("url parts : %T\n", partsOfUrl) // *url.URL
	anotherUrl := partsOfUrl.String()
	fmt.Println("another url : ", anotherUrl) // https://jsonplaceholder.typicode.com/comments?postId=1&userId=234

}
