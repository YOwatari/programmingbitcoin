package ecc

import "math/big"

func A() *S256Field {
	a, _ := NewS256Field(big.NewInt(0))
	return a
}

func B() *S256Field {
	b, _ := NewS256Field(big.NewInt(7))
	return b
}

func N() *big.Int {
	n, _ := new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	return n
}

func G() *S256Point {
	x, _ := new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	y, _ := new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	g, _ := NewS256PointFromBigInt(x, y)
	return g
}

type S256Point struct {
	*Point
}

func NewS256PointFromBigInt(x *big.Int, y *big.Int) (*S256Point, error) {
	X, err := NewS256Field(x)
	if err != nil {
		return nil, err
	}
	Y, err := NewS256Field(y)
	if err != nil {
		return nil, err
	}
	p, err := NewPoint(X, Y, A(), B())
	if err != nil {
		return nil, err
	}
	return &S256Point{p}, nil
}

func (p *S256Point) RMul(coef *big.Int) *S256Point {
	c := new(big.Int).Mod(coef, N())
	return &S256Point{p.Point.RMul(c)}
}
