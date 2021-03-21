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

func NewPoint(x FieldInterface, y FieldInterface, a FieldInterface, b FieldInterface) (*Point, error) {
	if x.IsInf() || y.IsInf() {
		return &Point{x, y, a, b, nil}, nil
	}

	verifyY, err := y.Copy().Pow(big.NewInt(2)).Calc()
	if err != nil {
		return nil, err
	}

	verifyX, err := x.Copy().Pow(big.NewInt(3)).Add(x.Copy().Mul(a)).Add(b).Calc()
	if err != nil {
		return nil, err
	}

	if verifyY.Ne(verifyX) {
		return nil, fmt.Errorf("(%#v, %#v) is not on the curve", x, y)
	}
	return &Point{X: x, Y: y, A: a, B: b, Err: nil}, nil
}

func (p *Point) String() string {
	if p.X.IsInf() {
		return "Point(infinity)"
	}
	return fmt.Sprintf("Point(%s, %s)_%s_%s", p.X, p.Y, p.A, p.B)
}

func (p *Point) Eq(other *Point) bool {
	if p.X.IsInf() || other.X.IsInf() {
		return p.X.IsInf() && other.X.IsInf()
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
	return &Point{p.X.Copy(), p.Y.Copy(), p.A.Copy(), p.B.Copy(), p.Err}
}

func (p *Point) Add(other *Point) *Point {
	if p.A.Ne(other.A) || p.B.Ne(other.B) {
		p.Err = fmt.Errorf("points %s, %s are not on the same curve", p, other)
		return p
	}

	if p.X.IsInf() {
		return other
	}

	if other.X.IsInf() {
		return p
	}

	if p.X.Eq(other.X) && p.Y.Ne(other.Y) {
		return p.Inf()
	}

	if p.Ne(other) {
		s := other.Y.Copy().Sub(p.Y).Div(other.X.Copy().Sub(p.X))
		x := s.Copy().Pow(big.NewInt(2)).Sub(p.X).Sub(other.X)
		y := s.Copy().Mul(p.X.Copy().Sub(x)).Sub(p.Y)
		*p = Point{x, y, p.A, p.B,p.Err}
		return p
	}

	zero := p.Y.Copy().RMul(big.NewInt(0))
	if p.Eq(other) && p.Y.Eq(zero) {
		return p.Inf()
	}

	s := p.X.Copy().Pow(big.NewInt(2)).RMul(big.NewInt(3)).Add(p.A).Div(p.Y.Copy().RMul(big.NewInt(2)))
	x := s.Copy().Pow(big.NewInt(2)).Sub(p.X.Copy().RMul(big.NewInt(2)))
	y := s.Copy().Mul(p.X.Copy().Sub(x)).Sub(p.Y)
	*p = Point{x, y, p.A, p.B, p.Err}
	return p
}

func (p *Point) RMul(n *big.Int) *Point {
	coefficient := n
	current := p
	result := p.Inf()

	for coefficient.Sign() > 0 {
		if coefficient.Bit(1) > 0 {
			result.Add(current)
		}
		current.Add(current)
		coefficient = new(big.Int).Rsh(coefficient, 1)
	}

	return result
}

func (p *Point) Inf() *Point {
	p.X.Inf()
	p.Y.Inf()
	return p
}
