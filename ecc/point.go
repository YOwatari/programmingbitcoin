package ecc

import (
	"fmt"
	"math"
)

type Point struct {
	A int
	B int
	X float64
	Y float64
	Err error
}

func NewPoint(x float64, y float64, a int , b int) (*Point, error) {
	p := &Point{a, b, math.Floor(x), math.Floor(y), nil}
	if math.IsNaN(p.X) && math.IsNaN(p.Y) {
		return p, nil
	}
	if int(math.Pow(p.Y, 2)) != int(math.Pow(p.X, 3)) + p.A * int(p.X) + p.B {
		return nil, fmt.Errorf("(%d, %d) is not on the curve", int(x), int(y))
	}
	return p, nil
}

func (p *Point) String() string {
	if math.IsNaN(p.X) {
		return "Point(infinity)"
	}
	return fmt.Sprintf("Point(%d, %d)_%d_%d", int(p.X), int(p.Y), p.A, p.B)
}

func (p *Point) Eq(other *Point) bool {
	if math.IsNaN(p.X) {
		return math.IsNaN(other.X)
	}
	return p.X == other.X && p.Y == other.Y && p.A == other.A && p.B == other.B
}

func (p *Point) Ne(other *Point) bool {
	return !p.Eq(other)
}

func (p *Point) Calc() (*Point, error) {
	return p, p.Err
}

func (p *Point) Add(other *Point) *Point {
	if p.A != other.A || p.B != other.B {
		p.Err = fmt.Errorf("points %s, %s are not on the same curve", p, other)
		return p
	}

	if math.IsNaN(p.X) {
		return other
	}
	if math.IsNaN(other.X) {
		return p
	}

	if p.X == other.X && p.Y != other.Y {
		p.X = math.NaN()
		p.Y = math.NaN()
		return p
	}

	if p.Ne(other) {
		s := (other.Y - p.Y) / (other.X - p.X)
		x := s * s - p.X - other.X
		y := s * (p.X - x) - p.Y
		return &Point{p.A, p.B, x, y, nil}
	}

	if p.Eq(other) && p.Y == 0 {
		return &Point{p.A, p.B, math.NaN(), math.NaN(), nil}
	}

	s := (3 * p.X * p.X + float64(p.A)) / (2 * p.Y)
	x := s * s - 2 * p.X
	y := s * (p.X - x) - p.Y
	return &Point{p.A, p.B, x, y, nil}
}
