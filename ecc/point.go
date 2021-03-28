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
	if x == nil || y == nil {
		return &Point{X: nil, Y: nil, A: a, B: b}, nil
	}

	left, right := y.Copy(), x.Copy()
	_, err := left.Pow(left, big.NewInt(2)).Calc()
	if err != nil {
		return nil, err
	}
	_, err = right.Pow(right, big.NewInt(3)).Add(right, x.Copy().Mul(x, a)).Add(right, b).Calc()
	if err != nil {
		return nil, err
	}

	if left.Ne(right) {
		return nil, fmt.Errorf("(%#v, %#v) is not on the curve", x, y)
	}
	return &Point{X: x, Y: y, A: a, B: b}, nil
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
	/*_, err := NewPoint(p.X, p.Y, p.A, p.B)
	if err != nil {
		return p, err
	}*/
	return p, p.Err
}

func (p *Point) Add(p1, p2 *Point) *Point {
	if p1.A.Ne(p2.A) || p1.B.Ne(p2.B) {
		*p = Point{
			X: p.X,
			Y: p.Y,
			A: p.A,
			B: p.B,
			Err: fmt.Errorf("points %s, %s are not on the same curve", p1, p2),
		}
		return p
	}
	a, b := p1.A, p1.B

	if p1.X == nil {
		*p = Point{X: p2.X, Y: p2.Y, A: a, B: b, Err: p.Err}
		return p
	}

	if p2.X == nil {
		*p = Point{X: p1.X, Y: p1.Y, A: a, B: b, Err: p.Err}
		return p
	}

	// p1.x == p2.x, p1.y != p2.y
	if p1.X.Eq(p2.X) && p1.Y.Ne(p2.Y) {
		*p = Point{X: nil, Y: nil, A: a, B: b, Err: p.Err}
		return p
	}

	x1, y1 := p1.X, p1.Y
	x2, y2 := p2.X, p2.Y

	// x1 != x2
	if x1.Ne(x2) {
		s := y2.Copy()
		s.Sub(y2, y1)
		tmp := x2.Copy()
		tmp.Sub(x2, x1)
		s.Div(s, tmp)
		x3 := s.Copy()
		x3.Pow(x3, big.NewInt(2)).Sub(x3, x1).Sub(x3, x2)
		y3 := s.Copy()
		y3.Mul(y3, tmp.Sub(x1, x3)).Sub(y3, y1)
		*p = Point{X: x3, Y: y3, A: a, B: b, Err: p.Err}
		return p
	}

	// p1 == p2, y == 0
	zero := x1.Copy()
	zero.RMul(zero, big.NewInt(0))
	if p1.Eq(p2) && p1.Y.Eq(zero) {
		*p = Point{X: nil, Y: nil, A: a, B: b, Err: p.Err}
		return p
	}

	// p1 == p2
	s := x1.Copy()
	s.Pow(x1, big.NewInt(2)).RMul(s, big.NewInt(3)).Add(s, a)
	tmp := y1.Copy()
	tmp.RMul(y1, big.NewInt(2))
	s.Div(s, tmp)
	x3 := s.Copy()
	x3.Pow(s, big.NewInt(2)).Sub(x3, tmp.RMul(x1, big.NewInt(2)))
	y3 := s.Copy()
	y3.Mul(s, tmp.Sub(x1, x3)).Sub(y3, y1)
	*p = Point{X: x3, Y: y3, A: a, B: b, Err: p.Err}
	return p
}

func (p *Point) RMul(r *Point, n *big.Int) *Point {
	coef := new(big.Int).Set(n)
	current := &Point{X: r.X, Y: r.Y, A: r.A, B: r.B, Err: r.Err}
	result := &Point{X: nil, Y: nil, A: r.A, B: r.B, Err: r.Err}

	for coef.Sign() > 0 {
		if coef.Bit(0) == 1 {
			result.Add(result, current)
		}
		current.Add(current, current)
		coef.Rsh(coef, 1)
	}

	*p = *result
	return p
}
