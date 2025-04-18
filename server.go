package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func getRequest(responseWriter http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(responseWriter, "<h1>Hello, World!</h1>")
}

func main() {
	http.HandleFunc("/", getRequest)
	http.ListenAndServe(PORT, nil)
}
