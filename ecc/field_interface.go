package ecc

import "math/big"

type FieldInterface interface {
	Eq(other FieldInterface) bool
	Ne(other FieldInterface) bool

	Calc() (FieldInterface, error)
	Copy() FieldInterface

	Add(a, b FieldInterface) FieldInterface
	Sub(a, b FieldInterface) FieldInterface
	Mul(a, b FieldInterface) FieldInterface
	Div(a, b FieldInterface) FieldInterface
	Pow(n FieldInterface, exp *big.Int) FieldInterface
	RMul(n FieldInterface, coef *big.Int) FieldInterface
}
