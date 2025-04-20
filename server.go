package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var LOG_DIR = filepath.Join(".", "logs")

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
	// make sure log dir exists
	errDir := os.MkdirAll(LOG_DIR, os.ModePerm)
	if errDir != nil {
		log.Fatal(errDir)
	}

	// start server
	http.HandleFunc("/log", dataLogger)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	errServer := http.ListenAndServe(PORT, nil)
	// errServer := http.ListenAndServeTLS(PORT, "server.crt", "server.key", nil)
	if errServer != nil {
		log.Fatal(errServer)
	}
}
