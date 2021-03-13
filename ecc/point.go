package ecc

import (
	"fmt"
)

type Point struct {
	X   FieldInterface
	Y   FieldInterface
	A   FieldInterface
	B   FieldInterface
	Err error
}

func NewPoint(x FieldInterface, y FieldInterface, a FieldInterface, b FieldInterface) (*Point, error) {
	if x == nil && y == nil {
		return &Point{nil, nil, a, b, nil}, nil
	}

	y2, err := y.Copy().Pow(2).Calc()
	if err != nil {
		return nil, err
	}
	x3, err := x.Copy().Pow(3).Add(x.Copy().Mul(a)).Add(b).Calc()
	if err != nil {
		return nil, err
	}

	if y2.Ne(x3) {
		return nil, fmt.Errorf("(%#v, %#v) is not on the curve", x, y)
	}
	return &Point{X: x, Y: y, A: a, B: b, Err: nil}, nil
}

func (p *Point) String() string {
	if p.X == nil {
		return "Point(infinity)"
	}
	return fmt.Sprintf("Point(%s, %s)_%s_%s", p.X, p.Y, p.A, p.B)
}

func (p *Point) Eq(other *Point) bool {
	if p.X == nil {
		return other.X == nil
	}
	return p.X.Eq(other.X) && p.Y.Eq(other.Y) && p.A.Eq(other.A) && p.B.Eq(other.B)
}

func (p *Point) Ne(other *Point) bool {
	return !p.Eq(other)
}

func (p *Point) Calc() (*Point, error) {
	return p, p.Err
}

func (p *Point) Copy() *Point {
	return &Point{p.X, p.Y, p.A, p.B, p.Err}
}

func (p *Point) Add(other *Point) *Point {
	if p.A.Ne(other.A) || p.B.Ne(other.B) {
		p.Err = fmt.Errorf("points %s, %s are not on the same curve", p, other)
		return p
	}

	if p.X == nil {
		return other
	}
	if other.X == nil {
		return p
	}

	if p.X.Eq(other.X) && p.Y.Ne(other.Y) {
		p.X = nil
		p.Y = nil
		return p
	}

	if p.Ne(other) {
		s := other.Y.Copy().Sub(p.Y).Div(other.X.Copy().Sub(p.X))
		x := s.Copy().Pow(2).Sub(p.X).Sub(other.X)
		y := s.Copy().Mul(p.X.Copy().Sub(x)).Sub(p.Y)
		*p = Point{x, y, p.A, p.B,nil}
		return p
	}

	zero := p.Y.Copy().MulInt(0)
	if p.Eq(other) && p.Y.Eq(zero) {
		*p = Point{nil, nil, p.A, p.B, nil}
		return p
	}

	s := p.X.Copy().Pow(2).MulInt(3).Add(p.A).Div(p.Y.Copy().MulInt(2))
	x := s.Copy().Pow(2).Sub(p.X.Copy().MulInt(2))
	y := s.Copy().Mul(p.X.Copy().Sub(x)).Sub(p.Y)
	*p = Point{x, y, p.A, p.B, nil}
	return p
}
