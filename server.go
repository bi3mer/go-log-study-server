package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const PORT = ":8080"

// const PORT = ":443"
func dataLogger(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	fmt.Printf("Received POST request with body: %s\n", body)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "POST request received successfully")
}

func main() {
	http.HandleFunc("/log", dataLogger)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	err := http.ListenAndServe(PORT, nil)
	// err := http.ListenAndServeTLS(PORT, "server.crt", "server.key", nil)
	if err != nil {
		log.Fatal(err)
	}
}
