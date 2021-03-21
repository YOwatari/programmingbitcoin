package ecc

import "math/big"

type FieldInterface interface {
	Eq(other FieldInterface) bool
	Ne(other FieldInterface) bool

	Calc() (FieldInterface, error)
	Copy() FieldInterface

	IsInf() bool
	Inf()

	Add(other FieldInterface) FieldInterface
	Sub(other FieldInterface) FieldInterface
	Mul(other FieldInterface) FieldInterface
	Div(other FieldInterface) FieldInterface
	Pow(exp *big.Int) FieldInterface
	RMul(coef *big.Int) FieldInterface
}
