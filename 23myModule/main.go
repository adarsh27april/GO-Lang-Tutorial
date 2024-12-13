package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("mod in golang")
	greeter()

	r := mux.NewRouter()
	r.HandleFunc("/", serveHome).Methods("GET") // only serving get method on route '/'
	// we are just passing ref. of serveHome method
	// i.e., on '/' route with GET Method serveHome method will be called.
	log.Fatal(http.ListenAndServe(":4000", r))
}
func greeter() {
	fmt.Println("welcome to go server")
}
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to golang series on YT</h1>"))
}
