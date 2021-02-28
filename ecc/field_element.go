package ecc

import (
	"fmt"
)

type FieldElement struct {
	Num int
	Prime int
}

func NewFieldElement(num int, prime int) (*FieldElement, error) {
	if num >= prime || num < 0 {
		return nil, fmt.Errorf("num %d not in field range 0 to %d", num, prime - 1)
	}
	return &FieldElement{
		Num: num,
		Prime: prime,
	}, nil
}

func (e *FieldElement) String() string {
	return fmt.Sprintf("FieldElement_%d(%d)", e.Prime, e.Num)
}

func (e *FieldElement) Eq(other *FieldElement) bool {
	return e.Num == other.Num && e.Prime == other.Prime
}

func (e *FieldElement) Ne(other *FieldElement) bool {
	return e.Num != other.Num || e.Prime != other.Prime
}

type CalcFunc func(*FieldElement) error

func (e *FieldElement) Calc(calcFuncs ...CalcFunc) error {
	for _, f := range calcFuncs {
		if err := f(e); err != nil {
			return err
		}
	}
	return nil
}

func Add(other *FieldElement) CalcFunc {
	return func(elm *FieldElement) error {
		if elm.Prime != other.Prime {
			return fmt.Errorf("cannot add two numbers in different Fields")
		}
		elm.Num = (elm.Num + other.Num) % elm.Prime
		return nil
	}
}

func Sub(other *FieldElement) CalcFunc {
	return func(elm *FieldElement) error {
		if elm.Prime != other.Prime {
			return fmt.Errorf("cannot sub two numbers in different Fields")
		}
		elm.Num = ((elm.Num - other.Num) % elm.Prime + elm.Prime) % elm.Prime
		return nil
	}
}

func Mul(other *FieldElement) CalcFunc {
	return func(elm *FieldElement) error {
		if elm.Prime != other.Prime {
			return fmt.Errorf("cannot multiply two numbers in different Fields")
		}
		elm.Num = (elm.Num * other.Num) % elm.Prime
		return nil
	}
}

func Pow(exponent int) CalcFunc {
	return func(elm *FieldElement) error {
		e := (exponent + (elm.Prime - 1)) % (elm.Prime - 1) // 0, p-2
		elm.Num = func(n int, e int, m int) int {
			p := 1
			for e > 0 {
				if e&1 == 1 {
					p = (p * n) % m
				}

				n = (n * n) % m
				if n == 1 {
					break
				}
				e >>= 1
			}
			return p
		}(elm.Num, e, elm.Prime)
		return nil
	}
}
