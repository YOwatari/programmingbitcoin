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

func (elm *FieldElement) Eq(other *FieldElement) bool {
	return elm.Num == other.Num && elm.Prime == other.Prime
}

func (elm *FieldElement) Ne(other *FieldElement) bool {
	return elm.Num != other.Num || elm.Prime != other.Prime
}

func (elm *FieldElement) Calc() (*FieldElement, error) {
	return elm, elm.Err
}

func (elm *FieldElement) Add(other *FieldElement) *FieldElement {
	if elm.Prime != other.Prime {
		elm.Err = fmt.Errorf("cannot add two numbers in different Fields")
		return elm
	}
	elm.Num = (elm.Num + other.Num) % elm.Prime
	return elm
}

func (elm *FieldElement) Sub(other *FieldElement) *FieldElement {
	if elm.Prime != other.Prime {
		elm.Err = fmt.Errorf("cannot sub two numbers in different Fields")
		return elm
	}
	elm.Num = ((elm.Num-other.Num)%elm.Prime + elm.Prime) % elm.Prime
	return elm
}

func (elm *FieldElement) Mul(other *FieldElement) *FieldElement {
	if elm.Prime != other.Prime {
		elm.Err = fmt.Errorf("cannot multiply two numbers in different Fields")
		return elm
	}
	elm.Num = (elm.Num * other.Num) % elm.Prime
	return elm
}

func (elm *FieldElement) Pow(exp int) *FieldElement {
	e := (exp + (elm.Prime - 1)) % (elm.Prime - 1) // 0, p-2
	elm.Num = func(n int, exp int, mod int) int {
		p := 1
		for exp > 0 {
			if exp&1 == 1 {
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

func (elm *FieldElement) Div(other *FieldElement) *FieldElement {
	if elm.Prime != other.Prime {
		elm.Err = fmt.Errorf("cannot division two numbers in different Fields")
		return elm
	}
	return elm.Mul(other.Pow(other.Prime - 2))
}
