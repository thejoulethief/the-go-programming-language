package main

import (
	"bufio"
	"fmt"
	"os"
)

// Get lines of text from command line, put them into a hash table i.e a map and print out duplicates
// Behold the power of hash tables!
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
	//fmt.Println(myMap)
	for key, value := range counts {
		fmt.Printf("%s:%d \n", key, value)
	}
}
