package main

import (
	"log"
	"net/http"
)

const PORT = ":8080"

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	err := http.ListenAndServe(PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}
