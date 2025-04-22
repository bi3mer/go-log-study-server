package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Post("http://127.0.0.1:8080/condition", "text/plain", nil)
	if err != nil {
		fmt.Print("hi0")
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print("hi")
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
