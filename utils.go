package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Pick a random number between the min and max received
func random(min, max int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(max-min) + min
}

// Handle any potential errors in the application
func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}
