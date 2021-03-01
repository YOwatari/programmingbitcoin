package ecc

import (
	"fmt"
	"math"
)

type Point struct {
	A int
	B int
	X int
	Y int
}

func NewPoint(x int, y int, a int , b int) (*Point, error) {
	if int(math.Pow(float64(y), 2)) != int(math.Pow(float64(x), 3)) + a * x + b {
		return nil, fmt.Errorf("(%d, %d) is not on the curve", x, y)
	}
	return &Point{a, b, x, y}, nil
}

func (p Point) String() string {
	return fmt.Sprintf("Point(%d, %d)_%d_%d", p.X, p.Y, p.A, p.B)
}

func (p *Point) Eq(other *Point) bool {
	return p.X == other.X && p.Y == other.Y && p.A == other.A && p.B == other.B
}

func (p *Point) Ne(other *Point) bool {
	return p.X != other.X || p.Y != other.Y || p.A != other.A || p.B != other.B
}
