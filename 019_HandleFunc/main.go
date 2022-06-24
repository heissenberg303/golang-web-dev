package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "dog dog dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "cat cat cat")
}

func main() {

	// instead of create a type hotdog implement Handler interface (Handler instance)
	// we can just create function that take ResponseWriter and pointer to Request
	// as parameters
	http.HandleFunc("/dog", d)
	http.HandleFunc("/cat", c)

	http.ListenAndServe(":8080", nil)
}
