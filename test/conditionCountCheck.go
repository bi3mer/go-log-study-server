package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

var M = sync.Mutex{}

func getCondition(counts map[string]int) {
	resp, err := http.Post("http://127.0.0.1:8080/condition", "text/plain", nil)
	if err != nil {
		fmt.Println(err)
		M.Lock()
		counts["error"]++
		M.Unlock()
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		M.Lock()
		counts["error"]++
		M.Unlock()
		return
	}

	M.Lock()
	counts[string(body)]++
	M.Unlock()
}

func main() {
	counts := map[string]int{}
	counts["random"] = 0
	counts["mean"] = 0
	counts["distance"] = 0
	counts["error"] = 0

	var wg sync.WaitGroup
	for range 200 {
		wg.Add(1)

		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
			getCondition(counts)
		}()
	}

	wg.Wait()

	println("random:   ", counts["random"])
	println("mean:     ", counts["mean"])
	println("distance: ", counts["distance"])
	println("error:    ", counts["error"])
}
