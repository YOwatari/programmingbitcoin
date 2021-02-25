package main

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

func (e FieldElement) Ne(other *FieldElement) bool {
	return e.Num != other.Num || e.Prime != other.Prime
}
