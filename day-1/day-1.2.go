package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := flag.String("filename", "input.txt", "input filename")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	seen := make(map[int]bool)
	freq := 0
	seen[freq] = true

	for {
		// start at the top so we can loop over and over
		file.Seek(0, io.SeekStart)
		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			line := scanner.Text()

			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}

			freq += num

			// check if our freq was already seen
			if seen[freq] {
				fmt.Println(freq)
				os.Exit(0)
			}

			seen[freq] = true
		}
	}
}
