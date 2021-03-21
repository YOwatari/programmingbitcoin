package exercise_test

import (
	"fmt"
	"github.com/YOwatari/programmingbitcoin/ecc"
)

func ExampleChapter3_one()  {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)

	points := []struct{
		x int64
		y int64
	} {
		{192, 105},
		{17, 56},
		{200, 119},
		{1, 193},
		{42, 99},
	}

	for _, p := range points {
		y, _ := ecc.NewFieldElementFromInt64(p.y, prime)
		x, _ := ecc.NewFieldElementFromInt64(p.x, prime)
		_, err := ecc.NewPoint(x, y, a, b)
		fmt.Printf("(%d, %d): %t\n", p.x, p.y, err == nil)
	}

	// Output:
	// (192, 105): true
	// (17, 56): true
	// (200, 119): false
	// (1, 193): true
	// (42, 99): false
}

func ExampleChapter3_two() {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)

	type point struct {
		x int64
		y int64
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
		y1, _ := ecc.NewFieldElementFromInt64(c.a.y, prime)
		x1, _ := ecc.NewFieldElementFromInt64(c.a.x, prime)
		y2, _ := ecc.NewFieldElementFromInt64(c.b.y, prime)
		x2, _ := ecc.NewFieldElementFromInt64(c.b.x, prime)
		p1, err := ecc.NewPoint(x1, y1, a, b)
		if err != nil {
			panic(err)
		}
		p2, err := ecc.NewPoint(x2, y2, a, b)
		if err != nil {
			panic(err)
		}

		p, err := p1.Copy().Add(p2).Calc()
		if err != nil {
			panic(err)
		}

		fmt.Printf("(%d, %d) + (%d, %d) = (%d, %d)\n", c.a.x, c.a.y, c.b.x, c.b.y, p.X.(*ecc.FieldElement).Num, p.Y.(*ecc.FieldElement).Num)
	}

	// Output:
	// (170, 142) + (60, 139) = (220, 181)
	// (47, 71) + (17, 56) = (215, 68)
	// (143, 98) + (76, 66) = (47, 71)
}

func ExampleChapter3_four() {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)

	cases := []struct{
		n int
		x int64
		y int64
	} {
		{
			2, 192, 105,
		},
		{
			2, 143, 98,
		},
		{
			2, 47, 71,
		},
		{
			4, 47, 71,
		},
		{
			8, 47, 71,
		},
		{
			21, 47, 71,
		},
	}

	for _, c := range cases {
		ansY, _ := ecc.NewFieldElementFromInt64(c.y, prime)
		ansX, _ := ecc.NewFieldElementFromInt64(c.x, prime)
		ansP, _ := ecc.NewPoint(ansX, ansY, a, b)
		for i := 0; i < c.n - 1; i++ {
			y, _ := ecc.NewFieldElementFromInt64(c.y, prime)
			x, _ := ecc.NewFieldElementFromInt64(c.x, prime)
			p, _ := ecc.NewPoint(x, y, a, b)
			ansP = ansP.Copy().Add(p)
		}
		p, err := ansP.Calc()
		if err != nil {
			panic(err)
		}

		if p.X.IsInf() || p.Y.IsInf() {
			fmt.Printf("%d * (%d, %d) = %s\n", c.n, c.x, c.y, p)
		} else {
			fmt.Printf("%d * (%d, %d) = Point(%d, %d)\n", c.n, c.x, c.y, p.X.(*ecc.FieldElement).Num, p.Y.(*ecc.FieldElement).Num)
		}
	}

	// Output:
	// 2 * (192, 105) = Point(49, 71)
	// 2 * (143, 98) = Point(64, 168)
	// 2 * (47, 71) = Point(36, 111)
	// 4 * (47, 71) = Point(194, 51)
	// 8 * (47, 71) = Point(116, 55)
	// 21 * (47, 71) = Point(infinity)
}

func ExampleChapter3_five()  {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)
	x, _ := ecc.NewFieldElementFromInt64(15, prime)
	y, _ := ecc.NewFieldElementFromInt64(86, prime)
	p, _ := ecc.NewPoint(x, y, a, b)

	ansP := p.Copy()
	inf := p.Copy().Inf()

	n := 1
	for ansP.Ne(inf) {
		ansP.Add(p)
		n++
	}

	_, err := ansP.Calc()
	if err != nil {
		panic(err)
	}

	fmt.Println(n)

	// Output:
	// 7
}
