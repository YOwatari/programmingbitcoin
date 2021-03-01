package ecc_test

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/YOwatari/programmingbitcoin/ecc"
)

func TestNewFieldElement(t *testing.T) {
	t.Run("Succeeds", func(t *testing.T) {
		actual, err := ecc.NewFieldElement(0, 11)
		if err != nil {
			t.Error(err)
		}
		expected, _ := ecc.NewFieldElement(0, 11)

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
			got, err := ecc.NewFieldElement(c.num, c.prime)
			if err == nil || got != nil {
				t.Error("should fail")
			}
		})
	}
}

func TestFieldElement_Eq(t *testing.T) {
	cases := []struct {
		a        *ecc.FieldElement
		b        *ecc.FieldElement
		expected bool
	}{
		{
			&ecc.FieldElement{Num: 7, Prime: 13},
			&ecc.FieldElement{Num: 7, Prime: 13},
			true,
		},
		{
			&ecc.FieldElement{Num: 7, Prime: 13},
			&ecc.FieldElement{Num: 6, Prime: 13},
			false,
		},
		{
			&ecc.FieldElement{Num: 7, Prime: 13},
			&ecc.FieldElement{Num: 7, Prime: 11},
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
		a        *ecc.FieldElement
		b        *ecc.FieldElement
		expected bool
	}{
		{
			&ecc.FieldElement{Num: 7, Prime: 13},
			&ecc.FieldElement{Num: 7, Prime: 13},
			false,
		},
		{
			&ecc.FieldElement{Num: 7, Prime: 13},
			&ecc.FieldElement{Num: 6, Prime: 13},
			true,
		},
		{
			&ecc.FieldElement{Num: 7, Prime: 13},
			&ecc.FieldElement{Num: 7, Prime: 11},
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

func TestFieldElement_Add(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		a, err := ecc.NewFieldElement(0, 1)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := a.Add(&ecc.FieldElement{Num: 0, Prime: 3}).Calc(); err == nil {
			t.Error("should fail to add two numbers in different Fields")
		}
	})

	cases := []struct {
		a   *ecc.FieldElement
		expected *ecc.FieldElement
	}{
		{
			(&ecc.FieldElement{Num: 7, Prime: 13}).Add(&ecc.FieldElement{Num: 12, Prime: 13}),
			&ecc.FieldElement{Num: 6, Prime: 13},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Eq(c.expected) != true {
				diff := cmp.Diff(actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Sub(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		a, err := ecc.NewFieldElement(0, 1)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := a.Sub(&ecc.FieldElement{Num: 0, Prime: 3}).Calc(); err == nil {
			t.Error("should fail to sub two numbers in different Fields")
		}
	})

	cases := []struct {
		a   *ecc.FieldElement
		expected *ecc.FieldElement
	}{
		{
			(&ecc.FieldElement{Num: 7, Prime: 13}).Sub(&ecc.FieldElement{Num: 6, Prime: 13}),
			&ecc.FieldElement{Num: 1, Prime: 13},
		},
		{
			(&ecc.FieldElement{Num: 7, Prime: 13}).Sub(&ecc.FieldElement{Num: 8, Prime: 13}),
			&ecc.FieldElement{Num: 12, Prime: 13},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Eq(c.expected) != true {
				diff := cmp.Diff(actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Mul(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		actual, err := ecc.NewFieldElement(0, 1)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := actual.Mul(&ecc.FieldElement{Num: 0, Prime: 3}).Calc(); err == nil {
			t.Error("should fail to multiply two numbers in different Fields")
		}
	})

	cases := []struct {
		a   *ecc.FieldElement
		expected *ecc.FieldElement
	}{
		{
			(&ecc.FieldElement{Num: 3, Prime: 13}).Mul(&ecc.FieldElement{Num: 12, Prime: 13}),
			&ecc.FieldElement{Num: 10, Prime: 13},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Eq(c.expected) != true {
				diff := cmp.Diff(actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Pow(t *testing.T) {
	cases := []struct {
		a   *ecc.FieldElement
		expected *ecc.FieldElement
	}{
		{
			(&ecc.FieldElement{Num: 3, Prime: 13}).Pow(3),
			&ecc.FieldElement{Num: 1, Prime: 13},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Eq(c.expected) != true {
				diff := cmp.Diff(actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}

func TestFieldElement_Div(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		a, err := ecc.NewFieldElement(0, 1)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := a.Div(&ecc.FieldElement{Num: 0, Prime: 3}).Calc(); err == nil {
			t.Error("should fail to division two numbers in different Fields")
		}
	})

	cases := []struct {
		a   *ecc.FieldElement
		expected *ecc.FieldElement
	}{
		{
			(&ecc.FieldElement{Num: 2, Prime: 19}).Div(&ecc.FieldElement{Num: 7, Prime: 19}),
			&ecc.FieldElement{Num: 3, Prime: 19},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Eq(c.expected) != true {
				diff := cmp.Diff(actual, c.expected)
				t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
			}
		})
	}
}
