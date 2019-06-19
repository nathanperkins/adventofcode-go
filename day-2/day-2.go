package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := flag.String("filename", "input.txt", "input filename")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	twos := 0
	threes := 0

	for scanner.Scan() {
		line := scanner.Text()
		count := make(map[rune]uint)

		for _, c := range line {
			count[c]++
		}

		for _, v := range count {
			if v == 2 {
				twos++
				break
			}
		}

		for _, v := range count {
			if v == 3 {
				threes++
				break
			}
		}
	}

	fmt.Printf("%v * %v = %v\n", twos, threes, twos*threes)
}
