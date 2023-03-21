package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

const WORKSIZE = 20000

func handleWork(res http.ResponseWriter, req *http.Request) {
	fmt.Println("Handling a new incoming Request!")
	// Check if the HTTP method is GET
	if req.Method != "GET" {
		http.Error(res, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	//Do the work
	total := 0
	for i := 0; i < WORKSIZE; i++ {
		total += rand.Intn(1000)
	}
	avg := total / WORKSIZE

	// Send a response with the submitted name
	fmt.Fprintf(res, "%d", avg)
}

func main() {
	// start of execution.

	// Set a new seed for the random number generator based on the current time
	rand.Seed(time.Now().UnixNano())

	// Listen on Server.
	http.HandleFunc("/", handleWork)
	http.ListenAndServe(":8080", nil)

}
