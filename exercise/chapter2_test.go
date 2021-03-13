package exercise_test

import (
	"fmt"
	"math"
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
		left := int(math.Pow(float64(p.y), 2))
		right := int(math.Pow(float64(p.x), 3)) + 5 * p.x + 7
		fmt.Printf("(%d, %d): %t\n", p.x, p.y, left == right)
	}

	// Output:
	// (2, 4): false
	// (-1, -1): true
	// (18, 77): true
	// (5, 7): false
}

func ExampleChapter2_four() {
	x1, y1 := 2, 5
	x2, y2 := -1, -1

	s := (y2 - y1) / (x2 -x1)
	x3 := int(math.Pow(float64(s), 2)) - x1 -x2
	y3 := s * (x1 - x3) - y1

	fmt.Printf("Point(%d, %d)_5_7 + Point(%d, %d)_5_7 = Point(%d, %d)_5_7", x1, y1, x2, y2, x3, y3)
	// Output:
	// Point(2, 5)_5_7 + Point(-1, -1)_5_7 = Point(3, -7)_5_7
}

func ExampleChapter2_six()  {
	x1, y1 := -1, -1
	x2, y2 := -1, -1

	s := (3 * int(math.Pow(float64(x1), 2)) + 5) / 2 * y1
	x3 := int(math.Pow(float64(s), 2)) - 2 * x1
	y3 := s * (x1 - x3) - y1

	fmt.Printf("Point(%d, %d)_5_7 + Point(%d, %d)_5_7 = Point(%d, %d)_5_7", x1, y1, x2, y2, x3, y3)
	// Output:
	// Point(-1, -1)_5_7 + Point(-1, -1)_5_7 = Point(18, 77)_5_7
}
