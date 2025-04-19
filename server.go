package main

import (
	"log"
	"net/http"
)

const PORT = ":8080"

// const PORT = ":443"

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	err := http.ListenAndServe(PORT, nil)
	// err := http.ListenAndServeTLS(PORT, "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
