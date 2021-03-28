package ecc

import (
	"fmt"
	"math/big"
)

type Signature struct {
	R *big.Int
	S *big.Int
}

func NewSignature(r, s *big.Int) *Signature {
	return &Signature{R: r, S: s}
}

func (s *Signature) String() string {
	return fmt.Sprintf("Signature(%064x, %064x)", s.R, s.S)
}
