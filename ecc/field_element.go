package ecc

import (
	"fmt"
)

type FieldElement struct {
	Num   int
	Prime int
	Err   error
}

func NewFieldElement(num int, prime int) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, fmt.Errorf("num %d not in field range 0 to %d", num, prime-1)
	}
	return &FieldElement{
		Num:   num,
		Prime: prime,
	}, nil
}

func (elm *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", elm.Prime, elm.Num)
}

func (elm *FieldElement) Eq(other FieldInterface) bool {
	elm2 := other.(*FieldElement)
	return elm.Num == elm2.Num && elm.Prime == elm2.Prime
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

func (elm *FieldElement) Add(other FieldInterface) FieldInterface {
	elm2 := other.(*FieldElement)
	if elm.Prime != elm2.Prime {
		elm.Err = fmt.Errorf("cannot add two numbers in different Fields")
		return elm
	}
	elm.Num = (elm.Num + elm2.Num) % elm.Prime
	return elm
}

func (elm *FieldElement) Sub(other FieldInterface) FieldInterface {
	elm2 := other.(*FieldElement)
	if elm.Prime != elm2.Prime {
		elm.Err = fmt.Errorf("cannot sub two numbers in different Fields")
		return elm
	}
	elm.Num = ((elm.Num-elm2.Num)%elm.Prime + elm.Prime) % elm.Prime
	return elm
}

func (elm *FieldElement) Mul(other FieldInterface) FieldInterface {
	elm2 := other.(*FieldElement)
	if elm.Prime != elm2.Prime {
		elm.Err = fmt.Errorf("cannot multiply two numbers in different Fields")
		return elm
	}
	elm.Num = (elm.Num * elm2.Num) % elm.Prime
	return elm
}

func (elm *FieldElement) Pow(exp int) FieldInterface {
	e := (exp + (elm.Prime - 1)) % (elm.Prime - 1) // 0, p-2
	elm.Num = func(n int, exp int, mod int) int {
		p := 1
		for exp > 0 {
			if exp & 1 == 1 {
				p = (p * n) % mod
			}

			n = (n * n) % mod
			if n == 1 {
				break
			}
			exp >>= 1
		}
		return p
	}(elm.Num, e, elm.Prime)
	return elm
}

func (elm *FieldElement) Div(other FieldInterface) FieldInterface {
	elm2 := other.(*FieldElement)
	if elm.Prime != elm2.Prime {
		elm.Err = fmt.Errorf("cannot division two numbers in different Fields")
		return elm
	}
	return elm.Mul(elm2.Pow(elm2.Prime - 2))
}

func (elm *FieldElement) RMul(c int) FieldInterface {
	elm.Num = elm.Num * c % elm.Prime
	return elm
}
