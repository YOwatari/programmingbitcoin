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
		_, err := ecc.NewPoint(float64(p.x), float64(p.y), 5, 7)
		fmt.Printf("(%d, %d): %t\n", p.x, p.y, err == nil)
	}

	// Output:
	// (2, 4): false
	// (-1, -1): true
	// (18, 77): true
	// (5, 7): false
}

func ExampleChapter2_four() {
	p1, _ := ecc.NewPoint(2, 5, 5, 7)
	p2, _ := ecc.NewPoint(-1, -1, 5, 7)
	p, _ := p1.Add(p2).Calc()
	fmt.Printf("%v + %v = %v", p1, p2, p)

	// Output:
	// Point(2, 5)_5_7 + Point(-1, -1)_5_7 = Point(3, -7)_5_7
}

func ExampleChapter2_six()  {
	p1, _ := ecc.NewPoint(-1, -1, 5, 7)
	p2, _ := ecc.NewPoint(-1, -1, 5, 7)
	p, _ := p1.Add(p2).Calc()
	fmt.Printf("%v + %v = %v", p1, p2, p)

	// Output:
	// Point(-1, -1)_5_7 + Point(-1, -1)_5_7 = Point(18, 77)_5_7
}
