package ecc

import (
	"fmt"
	"math/big"
)

type FieldElement struct {
	Num   *big.Int
	Prime *big.Int
	Err   error
	isInf bool
}

func NewFieldElementFromInt64(num int64, prime int64) (*FieldElement, error) {
	return NewFieldElement(big.NewInt(num), big.NewInt(prime))
}

func NewFieldElement(num *big.Int, prime *big.Int) (*FieldElement, error) {
	if num.Cmp(prime) >= 0 || num.Sign() < 0 {
		return nil, fmt.Errorf("num %d not in field range 0 to %d", num, new(big.Int).Div(prime, big.NewInt(1)))
	}
	return &FieldElement{Num: num, Prime: prime}, nil
}

func (elm *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", elm.Prime, elm.Num)
}

func (elm *FieldElement) Eq(other FieldInterface) bool {
	elm2 := other.(*FieldElement)
	return elm.Num.Cmp(elm2.Num) == 0 && elm.Prime.Cmp(elm2.Prime) == 0
}

func (elm *FieldElement) Ne(other FieldInterface) bool {
	return !elm.Eq(other)
}

func (elm *FieldElement) Calc() (FieldInterface, error) {
	return elm, elm.Err
}

func (elm *FieldElement) Copy() FieldInterface {
	return &FieldElement{Num: elm.Num, Prime: elm.Prime, Err: elm.Err, isInf: elm.isInf}
}

func (elm *FieldElement) IsInf() bool {
	return elm.isInf
}

func (elm *FieldElement) Inf() {
	elm.isInf = true
}

func (elm *FieldElement) Add(other FieldInterface) FieldInterface {
	elm2 := other.(*FieldElement)
	if elm.Prime.Cmp(elm2.Prime) != 0 {
		elm.Err = fmt.Errorf("cannot add two numbers in different Fields")
		return elm
	}

	n := new(big.Int)
	n.Add(elm.Num, elm2.Num).Mod(n, elm.Prime)

	elm.Num = n
	return elm
}

func (elm *FieldElement) Sub(other FieldInterface) FieldInterface {
	elm2 := other.(*FieldElement)
	if elm.Prime.Cmp(elm2.Prime) != 0 {
		elm.Err = fmt.Errorf("cannot sub two numbers in different Fields")
		return elm
	}

	n := new(big.Int)
	n.Sub(elm.Num, elm2.Num).Mod(n, elm.Prime)
	if n.Sign() < 0 {
		n.Add(n, elm.Prime)
	}

	elm.Num = n
	return elm
}

func (elm *FieldElement) Mul(other FieldInterface) FieldInterface {
	elm2 := other.(*FieldElement)
	if elm.Prime.Cmp(elm2.Prime) != 0 {
		elm.Err = fmt.Errorf("cannot multiply two numbers in different Fields")
		return elm
	}

	n := new(big.Int)
	n.Mul(elm.Num, elm2.Num).Mod(n, elm.Prime)

	elm.Num = n
	return elm
}

func (elm *FieldElement) Div(other FieldInterface) FieldInterface {
	elm2 := other.(*FieldElement)
	if elm.Prime.Cmp(elm2.Prime) != 0 {
		elm.Err = fmt.Errorf("cannot division two numbers in different Fields")
		return elm
	}

	m := new(big.Int).Sub(elm2.Prime, big.NewInt(2))
	return elm.Mul(elm2.Pow(m))
}

func (elm *FieldElement) Pow(exp *big.Int) FieldInterface {
	m := new(big.Int).Div(elm.Prime, big.NewInt(1))
	e := new(big.Int)
	e.Add(exp, m).Mod(e, m) // 0, p-2

	n := new(big.Int)
	n.Exp(elm.Num, e, elm.Prime)

	elm.Num = n
	return elm
}

func (elm *FieldElement) RMul(coef *big.Int) FieldInterface {
	n := new(big.Int)
	c := new(big.Int).Mod(coef, elm.Prime)
	n.Mul(elm.Num, c).Mod(n, elm.Prime)

	elm.Num = n
	return elm
}
