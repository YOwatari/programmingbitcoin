package ecc

import (
	"math/big"
)

var (
	G *S256Point
	_A *S256Field
	_B *S256Field
	N *big.Int
)

func init() {
	N, _ = new(big.Int).SetString("fffffffffffffffffffffffffffffffebaaedce6af48a03bbfd25e8cd0364141", 16)
	_A, _ = NewS256Field(big.NewInt(0))
	_B, _ = NewS256Field(big.NewInt(7))

	x, _ := new(big.Int).SetString("79be667ef9dcbbac55a06295ce870b07029bfcdb2dce28d959f2815b16f81798", 16)
	y, _ := new(big.Int).SetString("483ada7726a3c4655da4fbfc0e1108a8fd17b448a68554199c47d08ffb10d4b8", 16)
	G, _ = NewS256PointFromBigInt(x, y)
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
	p, err := NewPoint(X.FieldElement, Y.FieldElement, _A.FieldElement, _B.FieldElement)
	if err != nil {
		return nil, err
	}
	return &S256Point{p}, nil
}

func (p *S256Point) Add(p1, p2 *S256Point) *S256Point {
	result := new(Point).Add(p1.Point, p2.Point)
	*p = S256Point{result}
	return p
}

func (p *S256Point) RMul(r *S256Point, coef *big.Int) *S256Point {
	c := new(big.Int).Mod(coef, N)
	result := new(Point)
	result.RMul(r.Point, c)
	*p = S256Point{result}
	return p
}

func (p *S256Point) Verify(z *big.Int, sig *Signature) bool {
	invS := new(big.Int).Exp(sig.S, new(big.Int).Sub(N, big.NewInt(2)), N)
	u := new(big.Int)
	u.Mul(z, invS).Mod(u, N)
	v := new(big.Int)
	v.Mul(sig.R, invS).Mod(v, N)
	n := new(S256Point)
	n.Add(new(S256Point).RMul(G, u), new(S256Point).RMul(p, v))
	return n.X.(*FieldElement).Num.Cmp(sig.R) == 0
}
