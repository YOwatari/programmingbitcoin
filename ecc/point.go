package ecc

import (
	"fmt"
	"math/big"
)

type Point struct {
	X   FieldInterface
	Y   FieldInterface
	A   FieldInterface
	B   FieldInterface
	Err error
}

func NewPoint(rawX interface{}, rawY interface{}, a FieldInterface, b FieldInterface) (*Point, error) {
	if rawX == nil || rawY == nil {
		return &Point{nil, nil, a, b, nil}, nil
	}
	x, ok := rawX.(FieldInterface)
	if !ok {
		return nil, fmt.Errorf("rawX interface conversion")
	}
	y, ok := rawY.(FieldInterface)
	if !ok {
		return nil, fmt.Errorf("rawY interface conversion")
	}

	verifyY, err := y.Copy().Pow(2).Calc()
	if err != nil {
		return nil, err
	}

	verifyX, err := x.Copy().Pow(3).Add(x.Copy().Mul(a)).Add(b).Calc()
	if err != nil {
		return nil, err
	}

	if verifyY.Ne(verifyX) {
		return nil, fmt.Errorf("(%#v, %#v) is not on the curve", rawX, rawY)
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
	if other.X == nil {
		return p.X == nil
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
		*p = Point{x, y, p.A, p.B,p.Err}
		return p
	}

	zero := p.Y.Copy().RMul(0)
	if p.Eq(other) && p.Y.Eq(zero) {
		*p = Point{nil, nil, p.A, p.B, p.Err}
		return p
	}

	s := p.X.Copy().Pow(2).RMul(3).Add(p.A).Div(p.Y.Copy().RMul(2))
	x := s.Copy().Pow(2).Sub(p.X.Copy().RMul(2))
	y := s.Copy().Mul(p.X.Copy().Sub(x)).Sub(p.Y)
	*p = Point{x, y, p.A, p.B, p.Err}
	return p
}

func (p *Point) RMul(n *big.Int) *Point {
	coefficient := n
	current := p
	result := &Point{nil, nil, p.A, p.B, p.Err}

	for coefficient.Sign() > 0 {
		if coefficient.Bit(1) > 0 {
			result.Add(current)
		}
		current.Add(current)
		coefficient = new(big.Int).Rsh(coefficient, 1)
	}

	return result
}
