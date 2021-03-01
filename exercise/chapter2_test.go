package exercise_test

import (
	"fmt"

	"github.com/YOwatari/programmingbitcoin/ecc"
)

func ExampleChapter2_one()  {
	points := []struct{
		x int
		y int
	} {
		{2, 4 },
		{ -1, -1},
		{ 18, 77},
		{5, 7},
	}
	for _, p := range points {
		_, err := ecc.NewPoint(p.x, p.y, 5, 7)
		fmt.Printf("(%d, %d): %t\n", p.x, p.y, err == nil)
	}

	// Output:
	// (2, 4): false
	// (-1, -1): true
	// (18, 77): true
	// (5, 7): false
}
