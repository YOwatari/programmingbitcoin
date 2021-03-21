package ecc

type FieldInterface interface {
	Eq(other FieldInterface) bool
	Ne(other FieldInterface) bool

	Calc() (FieldInterface, error)
	Copy() FieldInterface

	Add(other FieldInterface) FieldInterface
	Sub(other FieldInterface) FieldInterface
	Mul(other FieldInterface) FieldInterface
	Pow(exp int) FieldInterface
	Div(other FieldInterface) FieldInterface
	RMul(c int) FieldInterface
}
