package main

import (
	"bufio"
	"fmt"
	"os"
)

// Get lines of text from command line, put them into a hash table i.e a map and print out duplicates
// If cmd line arg with filename is provided, read from that file line by line
// Behold the power of hash tables and O(1) lookup!

func main() {
	if len(os.Args) < 2 {
		findDuplicates(os.Stdin)
	} else {
		filename := os.Args[1]
		file, err := os.Open(filename)
		if err != nil {
			fmt.Printf("Error: %v", err)
		}
		findDuplicates(file)
	}
}

func findDuplicates(file *os.File) {
	counts := make(map[string]int)
	input := bufio.NewScanner(file)

	for input.Scan() {
		if input.Text() == "" {
			break
		}
		counts[input.Text()]++
	}
	for key, value := range counts {
		fmt.Printf("%s:%d \n", key, value)
	}
}
