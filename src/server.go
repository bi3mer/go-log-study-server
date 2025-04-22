package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

// //////////////////////////////////////////////////////////////
// Globals
const PORT = ":8080" // :443 for ssl

var LOG_DIR = filepath.Join(".", "logs")
var CONDITIONS = []string{"random", "mean", "distance"}
var BLOCK = blockNew(2, CONDITIONS)

// //////////////////////////////////////////////////////////////
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

	for {
		logID := uuid.New()
		path := filepath.Join(LOG_DIR, logID.String()+".json")
		_, err := os.Stat(path)

		if os.IsNotExist(err) {
			file, createFileErr := os.Create(path)
			if createFileErr != nil {
				http.Error(w, "Could not create a file to log to.", http.StatusInternalServerError)
				return
			}

			defer file.Close()

			file.Write(body)

			fmt.Printf("Data logged to '%s'\n", path)
			w.WriteHeader(http.StatusCreated)
			break
		}
	}
}

// Post request - get study condition
func getCondition(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Println("Condition request was not a POST request.")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Println("Received condition request.")

	w.WriteHeader(http.StatusCreated)
	fmt.Fprint(w, *blockGetCondition(BLOCK))
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

	// https://coe.northeastern.edu/computer/general-resources/vpn/
	// https://letsencrypt.org/
	// errServer := http.ListenAndServeTLS(PORT, "server.crt", "server.key", nil)
	if errServer != nil {
		log.Fatal(errServer)
	}
}
