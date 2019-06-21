package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	maxWidth  = 1000
	maxHeight = 1000
)

func main() {
	filename := flag.String("filename", "input.txt", "input filename")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil {
		log.Fatal(err)
	}

	grid := make(Grid, 0)
	claims := make([]Claim, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		newClaim := newClaim(scanner.Text())
		claims = append(claims, newClaim)
		grid.addClaim(&newClaim)
	}

	// fmt.Println(grid)
}

// Claim will have an ID and consist of a single rectangle
// with edges parallel to the edges of the fabric.
//
// LeftMargin: The number of inches between the left edge of the fabric and the left edge of the rectangle.
// TopMargin:  The number of inches between the top edge of the fabric and the top edge of the rectangle.
// Width:      The width of the rectangle in inches.
// Height:     The height of the rectangle in inches.
type Claim struct {
	ID, LeftMargin, TopMargin, Width, Height int
	FoundOverlap                             bool
}

// Grid holds a represesntation of the total fabric for the problem
type Grid map[int]map[int][]*Claim

// NewClaim creates a new claim from a string.
func newClaim(str string) Claim {
	var newClaim Claim

	fmt.Sscanf(str, "#%v @ %v,%v: %vx%v\n",
		&newClaim.ID,
		&newClaim.LeftMargin, &newClaim.TopMargin,
		&newClaim.Width, &newClaim.Height)

	return newClaim
}

// AddClaim updates a grade to reflect the
func (grid Grid) addClaim(claim *Claim) {
	for y := claim.TopMargin; y < claim.TopMargin+claim.Height; y++ {
		_, ok := grid[y]

		if !ok {
			grid[y] = make(map[int][]*Claim, 0)
		}

		for x := claim.LeftMargin; x < claim.LeftMargin+claim.Width; x++ {
			_, ok := grid[y][x]

			if !ok {
				grid[y][x] = make([]*Claim, 0)
			}

			fmt.Println("checking", x, y)

			for _, otherClaim := range grid[y][x] {
				claim.FoundOverlap = true
				otherClaim.FoundOverlap = true
			}

			grid[y][x] = append(grid[y][x], claim)
		}
	}
}
