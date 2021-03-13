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
		y := _pow(p.y, 2, prime)
		x := (_pow(p.x, 3, prime) + a * p.x + b) % prime
		fmt.Printf("(%d, %d): %t\n", p.x, p.y, y == x)
	}

	// Output:
	// (192, 105): true
	// (17, 56): true
	// (200, 119): false
	// (1, 193): true
	// (42, 99): false
}

func ExampleChapter3_two() {
	prime := 223
	a := 0
	b := 7

	type point struct {
		x int
		y int
	}
	cases := []struct{
		a point
		b point
	} {
		{
			point{170, 142},
			point{60, 139},
		},
		{
			point{47, 71},
			point{17, 56},
		},
		{
			point{143, 98},
			point{76, 66},
		},
	}

	for _, c := range cases {
		s := _div(_sub(c.b.y, c.a.y, prime), _sub(c.b.x, c.a.x, prime), prime)
		x3 := _sub(_sub(_pow(s, 2, prime), c.a.x, prime), c.b.x, prime)
		y3 := _sub(s * _sub(c.a.x, x3, prime), c.a.y, prime)

		y := _pow(y3, 2, prime)
		x := (_pow(x3, 3, prime) + a * x3 + b) % prime
		if y == x {
			fmt.Printf("(%d, %d) + (%d, %d) = (%d, %d)\n", c.a.x, c.a.y, c.b.x, c.b.y, x3, y3)
		}
	}

	// Output:
	// (170, 142) + (60, 139) = (220, 181)
	// (47, 71) + (17, 56) = (215, 68)
	// (143, 98) + (76, 66) = (47, 71)
}

func _div(a int, b int, mod int) int {
	return (a * _pow(b, mod - 2, mod)) % mod
}
