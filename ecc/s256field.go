package ecc

import (
	"fmt"
	"math/big"
)

func Prime() *big.Int {
	prime := new(big.Int)
	t256 := new(big.Int).Exp(big.NewInt(2), big.NewInt(256), nil)
	t32 := new(big.Int).Exp(big.NewInt(2), big.NewInt(32), nil)
	prime.Sub(t256, t32).Sub(prime, big.NewInt(977))
	return prime
}

type S256Field struct {
	*FieldElement
}

func NewS256Field(num *big.Int) (*S256Field, error) {
	f, err := NewFieldElement(num, Prime())
	if err != nil {
		return nil, err
	}
	return &S256Field{f}, nil
}

func (f *S256Field) String() string {
	return fmt.Sprintf("%064d", f.Num)
}
