package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func matchingChars(lineA string, lineB string) string {
	// return a string of all the chars that match between the lines
	// ex: matchingChars("abc", "xbc") would return "bc"

	// lines should be the same size for our purposes
	if len(lineA) != len(lineB) {
		fmt.Println(lineA, lineB)
		log.Fatal("Lines are not the same size!")
	}

	// build our string
	ret := ""
	for i := 0; i < len(lineA); i++ {
		if lineA[i] == lineB[i] {
			ret += string(lineA[i])
		}
	}

	return ret
}

func main() {
	filename := flag.String("filename", "input.txt", "input filename")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	lines := make([]string, 0)

	// load all the lines to memory
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// go through each possible pair of lines
	for a, lineA := range lines {
		for _, lineB := range lines[a+1:] {

			chars := matchingChars(lineA, lineB)
			diff := len(lineA) - len(chars)

			// check if our pair is good :)
			if diff == 1 {
				fmt.Println("Found pair:   ", lineA, lineB)
				fmt.Println("Common chars: ", chars)
				os.Exit(0)
			}
		}
	}

	fmt.Println("No match found.")
}
