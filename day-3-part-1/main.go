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

	scanner := bufio.NewScanner(file)
	grid := Grid{}

	for scanner.Scan() {
		newClaim := Claim{}

		// ex: #1 @ 916,616: 21x2
		fmt.Sscanf(scanner.Text(), "#%v @ %v,%v: %vx%v\n",
			&newClaim.ID,
			&newClaim.LeftMargin, &newClaim.TopMargin,
			&newClaim.Width, &newClaim.Height)

		grid.addClaim(newClaim)
	}

	fmt.Println(grid.inchesClaimedTwice())
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
}

// Grid is a 2D array used to hold the number of times each square inch is claimed.
type Grid [maxHeight][maxWidth]int

// addClaim will add one to each square on the grid that this claim claims.
func (grid *Grid) addClaim(claim Claim) {
	for x := 0; x < claim.Width; x++ {
		for y := 0; y < claim.Height; y++ {
			grid[claim.TopMargin+y][claim.LeftMargin+x]++
		}
	}
}

// inchesClaimedTwice will return the number of square inches on the grid
// that were claimed by at least two claims
func (grid *Grid) inchesClaimedTwice() int {
	ret := 0

	for x := 0; x < maxWidth; x++ {
		for y := 0; y < maxHeight; y++ {
			if grid[y][x] > 1 {
				ret++
			}
		}
	}

	return ret
}
