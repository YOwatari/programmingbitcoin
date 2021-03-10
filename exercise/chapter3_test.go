package exercise_test

import (
	"fmt"
)

func ExampleChapter3_one()  {
	prime := 223
	a := 0
	b := 7
	points := []struct{
		x int
		y int
	} {
		{192, 105},
		{17, 56},
		{200, 119},
		{1, 193},
		{42, 99},
	}
	for _, p := range points {
		y2 := (p.y * p.y) % prime
		x3 := ((p.x * p.x * p.x) + a * p.x + b) % prime
		fmt.Printf("(%d, %d): %t\n", p.x, p.y, y2 == x3)
	}

	// Output:
	// (192, 105): true
	// (17, 56): true
	// (200, 119): false
	// (1, 193): true
	// (42, 99): false
}
