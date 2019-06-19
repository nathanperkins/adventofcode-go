package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inputFilename := flag.String("filename", "input.txt",
		"input filename")
	flag.Parse()

	inputFile, err := os.Open(*inputFilename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		sum += n
	}

	fmt.Println(sum)
}
