package ecc

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFieldElement(t *testing.T) {
	t.Run("Succeeds", func(t *testing.T) {
		actual, err := NewFieldElement(0, 11)
		if err != nil {
			t.Error(err)
		}
		expected := &FieldElement{0, 11}

		if diff := cmp.Diff(actual, expected); diff != "" {
			t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
		}
	})

	cases := []struct {
		num   int
		prime int
	}{
		{
			10,
			1,
		},
		{
			-1,
			1,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("Fails_%d", i), func(t *testing.T) {
			got, err := NewFieldElement(c.num, c.prime)
			if err == nil || got != nil {
				t.Error("should fail")
			}
		})
	}
}

func TestFieldElement_Eq(t *testing.T) {
	cases := []struct {
		a        *FieldElement
		b        *FieldElement
		expected bool
	}{
		{
			&FieldElement{7, 13},
			&FieldElement{7, 13},
			true,
		},
		{
			&FieldElement{7, 13},
			&FieldElement{6, 13},
			false,
		},
		{
			&FieldElement{7, 13},
			&FieldElement{7, 11},
			false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if c.a.Eq(c.b) != c.expected {
				t.Errorf("FieldElement.Eq: %#v, %#v, expected: %t", c.a, c.b, c.expected)
			}
		})
	}
}

func TestFieldElement_Ne(t *testing.T) {
	cases := []struct {
		a        *FieldElement
		b        *FieldElement
		expected bool
	}{
		{
			&FieldElement{7, 13},
			&FieldElement{7, 13},
			false,
		},
		{
			&FieldElement{7, 13},
			&FieldElement{6, 13},
			true,
		},
		{
			&FieldElement{7, 13},
			&FieldElement{7, 11},
			true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if c.a.Ne(c.b) != c.expected {
				t.Errorf("FieldElement.Ne: %#v, %#v, expected: %t", c.a, c.b, c.expected)
			}
		})
	}
}

func TestAdd(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		actual := &FieldElement{0, 1}
		if err := actual.Calc(Add(&FieldElement{0, 3})); err == nil {
			t.Error("should fail to add two numbers in different Fields")
		}
	})

	cases := []struct {
		actual   *FieldElement
		cals     []CalcFieldElementFunc
		expected *FieldElement
	}{
		{
			&FieldElement{7, 13},
			[]CalcFieldElementFunc{Add(&FieldElement{12, 13})},
			&FieldElement{6, 13},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if err := c.actual.Calc(c.cals...); err != nil {
				t.Fatal(err)
			}

			if c.actual.Eq(c.expected) != true {
				diff := cmp.Diff(c.actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Sub(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		actual := &FieldElement{0, 1}
		if err := actual.Calc(Sub(&FieldElement{0, 3})); err == nil {
			t.Error("should fail to sub two numbers in different Fields")
		}
	})

	cases := []struct {
		actual   *FieldElement
		cals     []CalcFieldElementFunc
		expected *FieldElement
	}{
		{
			&FieldElement{7, 13},
			[]CalcFieldElementFunc{Sub(&FieldElement{6, 13})},
			&FieldElement{1, 13},
		},
		{
			&FieldElement{7, 13},
			[]CalcFieldElementFunc{Sub(&FieldElement{8, 13})},
			&FieldElement{12, 13},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if err := c.actual.Calc(c.cals...); err != nil {
				t.Fatal(err)
			}

			if c.actual.Eq(c.expected) != true {
				diff := cmp.Diff(c.actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestMul(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		actual := &FieldElement{0, 1}
		if err := actual.Calc(Mul(&FieldElement{0, 3})); err == nil {
			t.Error("should fail to multiply two numbers in different Fields")
		}
	})

	cases := []struct {
		actual   *FieldElement
		cals     []CalcFieldElementFunc
		expected *FieldElement
	}{
		{
			&FieldElement{3, 13},
			[]CalcFieldElementFunc{Mul(&FieldElement{12, 13})},
			&FieldElement{10, 13},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if err := c.actual.Calc(c.cals...); err != nil {
				t.Fatal(err)
			}

			if c.actual.Eq(c.expected) != true {
				diff := cmp.Diff(c.actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestPow(t *testing.T) {
	cases := []struct {
		actual   *FieldElement
		cals     []CalcFieldElementFunc
		expected *FieldElement
	}{
		{
			&FieldElement{3, 13},
			[]CalcFieldElementFunc{Pow(3)},
			&FieldElement{1, 13},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if err := c.actual.Calc(c.cals...); err != nil {
				t.Fatal(err)
			}

			if c.actual.Eq(c.expected) != true {
				diff := cmp.Diff(c.actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		actual := &FieldElement{0, 1}
		if err := actual.Calc(Div(&FieldElement{0, 3})); err == nil {
			t.Error("should fail to division two numbers in different Fields")
		}
	})

	cases := []struct {
		actual   *FieldElement
		cals     []CalcFieldElementFunc
		expected *FieldElement
	}{
		{
			&FieldElement{2, 19},
			[]CalcFieldElementFunc{Div(&FieldElement{7, 19})},
			&FieldElement{3, 19},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if err := c.actual.Calc(c.cals...); err != nil {
				t.Fatal(err)
			}

			if c.actual.Eq(c.expected) != true {
				diff := cmp.Diff(c.actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}
