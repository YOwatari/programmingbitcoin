package ecc

import (
	"fmt"
	"math/big"
)

type FieldElement struct {
	Num   *big.Int
	Prime *big.Int
	Err   error
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
	if other == nil {
		return false
	}
	o := other.(*FieldElement)
	return elm.Num.Cmp(o.Num) == 0 && elm.Prime.Cmp(o.Prime) == 0
}

func (elm *FieldElement) Ne(other FieldInterface) bool {
	return !elm.Eq(other)
}

func (elm *FieldElement) Calc() (FieldInterface, error) {
	return elm, elm.Err
}

func (elm *FieldElement) Copy() FieldInterface {
	return &FieldElement{Num: elm.Num, Prime: elm.Prime, Err: elm.Err}
}

func (elm *FieldElement) Add(a, b FieldInterface) FieldInterface {
	x, y := a.(*FieldElement), b.(*FieldElement)
	if x.Prime.Cmp(y.Prime) != 0 {
		*elm = FieldElement{
			Num: elm.Num,
			Prime: elm.Prime,
			Err: fmt.Errorf("cannot add two numbers in different Fields"),
		}
		return elm
	}

	num := new(big.Int)
	num.Add(x.Num, y.Num).Mod(num, x.Prime)

	*elm = FieldElement{Num: num, Prime: x.Prime, Err: elm.Err}
	return elm
}

func (elm *FieldElement) Sub(a, b FieldInterface) FieldInterface {
	x, y := a.(*FieldElement), b.(*FieldElement)
	if x.Prime.Cmp(y.Prime) != 0 {
		*elm = FieldElement{
			Num: elm.Num,
			Prime: elm.Prime,
			Err: fmt.Errorf("cannot sub two numbers in different Fields"),
		}
		return elm
	}

	num := new(big.Int)
	num.Sub(x.Num, y.Num).Mod(num, x.Prime)
	if num.Sign() < 0 {
		num.Add(num, x.Prime)
	}

	*elm = FieldElement{Num: num, Prime: x.Prime, Err: elm.Err}
	return elm
}

func (elm *FieldElement) Mul(a, b FieldInterface) FieldInterface {
	x, y := a.(*FieldElement), b.(*FieldElement)
	if x.Prime.Cmp(y.Prime) != 0 {
		*elm = FieldElement{
			Num: elm.Num,
			Prime: elm.Prime,
			Err: fmt.Errorf("cannot mul two numbers in different Fields"),
		}
		return elm
	}

	num := new(big.Int)
	num.Mul(x.Num, y.Num).Mod(num, x.Prime)

	*elm = FieldElement{Num: num, Prime: x.Prime, Err: elm.Err}
	return elm
}

func (elm *FieldElement) Div(a, b FieldInterface) FieldInterface {
	x, y := a.(*FieldElement), b.(*FieldElement)
	if x.Prime.Cmp(y.Prime) != 0 {
		*elm = FieldElement{
			Num:   elm.Num,
			Prime: elm.Prime,
			Err: fmt.Errorf("cannot division two numbers in different Fields"),
		}
		return elm
	}

	m := new(big.Int)
	m.Sub(x.Prime, big.NewInt(2))
	return x.Mul(x, y.Pow(y, m))
}

func (elm *FieldElement) Pow(n FieldInterface, exp *big.Int) FieldInterface {
	x := n.(*FieldElement)

	m := new(big.Int)
	e := new(big.Int)
	m.Sub(x.Prime, big.NewInt(1))
	e.Mod(exp, m)

	num := new(big.Int)
	num.Exp(x.Num, e, x.Prime)

	*elm = FieldElement{Num: num, Prime: x.Prime, Err: elm.Err}
	return elm
}

func (elm *FieldElement) RMul(n FieldInterface, coef *big.Int) FieldInterface {
	x := n.(*FieldElement)
	c := new(big.Int).Mod(coef, x.Prime)
	num := new(big.Int)
	num.Mul(x.Num, c).Mod(num, x.Prime)

	*elm = FieldElement{Num: num, Prime: x.Prime, Err: elm.Err}
	return elm
}
