package main

import (
	"fmt"
	"strings"
)

func main() {
	var words string = "selamat malam"

	wordsSlice := strings.Split(words, "")

	for i := 0; i < len(wordsSlice); i++ {
		fmt.Println(wordsSlice[i])
	}

	// Creating a map to hold character counts
	charCount := make(map[string]int)

	// Iterating over the slice to populate the map
	for _, char := range wordsSlice {
		// Incrementing the count of each character
		charCount[char]++
	}

	fmt.Println(charCount)
}
