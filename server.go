package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// ////////////////////////////////////////////////////////////////
// Globals
const PORT = ":8080" // :443 for ssl

var LOG_DIR = filepath.Join(".", "logs")

// ////////////////////////////////////////////////////////////////
// Post request - store data file
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
}

// Post request - get condition
func getCondition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, "POST request received successfully")
}

func main() {
	////////////////////////////////////////////////////////////////////
	// Make sure log dir exists
	errDir := os.MkdirAll(LOG_DIR, os.ModePerm)
	if errDir != nil {
		log.Fatal(errDir)
	}

	////////////////////////////////////////////////////////////////////
	// Make sure static directory exists before starting the server
	dirStatic, errStatic := os.Stat("./static")
	if os.IsNotExist(errStatic) {
		log.Fatal("Static directory must exist", errStatic)
	}

	if !dirStatic.IsDir() {
		log.Fatal("Found file named 'static'. Should be a directory.")
	}

	////////////////////////////////////////////////////////////////////
	// Start server
	http.HandleFunc("/log", dataLogger)
	http.HandleFunc("/condition", getCondition)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	errServer := http.ListenAndServe(PORT, nil)
	// errServer := http.ListenAndServeTLS(PORT, "server.crt", "server.key", nil)
	if errServer != nil {
		log.Fatal(errServer)
	}
}
