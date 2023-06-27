package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func generateRandomNumberString(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Create a byte slice to store the random digits
	digits := make([]byte, length)

	// Generate random digits
	for i := 0; i < length; i++ {
		digits[i] = byte(rand.Intn(10)) + '0' // Generate a random digit from 0 to 9
	}

	// Convert the byte slice to a string
	return string(digits)
}

func main() {
	randomNumber := generateRandomNumberString(8)
	fmt.Println(randomNumber)
}
