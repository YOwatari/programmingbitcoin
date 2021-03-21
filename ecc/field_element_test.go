package ecc_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/YOwatari/programmingbitcoin/ecc"
)

func TestNewFieldElement(t *testing.T) {
	t.Run("Succeeds", func(t *testing.T) {
		_, err := ecc.NewFieldElementFromInt64(0, 11)
		if err != nil {
			t.Error(err)
		}
	})

	cases := []struct {
		num   int64
		prime int64
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
			got, err := ecc.NewFieldElementFromInt64(c.num, c.prime)
			if err == nil || got != nil {
				t.Error("should fail")
			}
		})
	}
}

func _newFieldElement(num int64, prime int64) *ecc.FieldElement {
	elm, _ := ecc.NewFieldElementFromInt64(num, prime)
	return elm
}

func TestFieldElement_Eq(t *testing.T) {
	cases := []struct {
		a        *ecc.FieldElement
		b        *ecc.FieldElement
		expected bool
	}{
		{
			_newFieldElement(7, 13),
			_newFieldElement(7, 13),
			true,
		},
		{
			_newFieldElement(7, 13),
			_newFieldElement(6, 13),
			false,
		},
		{
			_newFieldElement(7, 13),
			_newFieldElement(7, 11),
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
			_newFieldElement(7, 13),
			_newFieldElement(7, 13),
			false,
		},
		{
			_newFieldElement(7, 13),
			_newFieldElement(6, 13),
			true,
		},
		{
			_newFieldElement(7, 13),
			_newFieldElement(7, 11),
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
		a, err := ecc.NewFieldElementFromInt64(0, 1)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := a.Add(_newFieldElement(0, 3)).Calc(); err == nil {
			t.Error("should fail to add two numbers in different Fields")
		}
	})

	cases := []struct {
		a   ecc.FieldInterface
		expected *ecc.FieldElement
	}{
		{
			_newFieldElement(7, 13).Add(_newFieldElement(12, 13)),
			_newFieldElement(6, 13),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Ne(c.expected) {
				t.Errorf("\n got: %s\n want: %s\n", actual, c.expected)
			}
		})
	}
}

func TestFieldElement_Sub(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		a, err := ecc.NewFieldElementFromInt64(0, 1)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := a.Sub(_newFieldElement(0, 3)).Calc(); err == nil {
			t.Error("should fail to sub two numbers in different Fields")
		}
	})

	cases := []struct {
		a   ecc.FieldInterface
		expected *ecc.FieldElement
	}{
		{
			_newFieldElement(7, 13).Sub(_newFieldElement(6, 13)),
			_newFieldElement(1, 13),
		},
		{
			_newFieldElement(7, 13).Sub(_newFieldElement(8, 13)),
			_newFieldElement(12, 13),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Ne(c.expected) {
				t.Errorf("\n got: %s\n want: %s\n", actual, c.expected)
			}
		})
	}
}

func TestFieldElement_Mul(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		actual, err := ecc.NewFieldElementFromInt64(0, 1)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := actual.Mul(_newFieldElement(0, 3)).Calc(); err == nil {
			t.Error("should fail to multiply two numbers in different Fields")
		}
	})

	cases := []struct {
		a   ecc.FieldInterface
		expected *ecc.FieldElement
	}{
		{
			_newFieldElement(3, 13).Mul(_newFieldElement(12, 13)),
			_newFieldElement(10, 13),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Ne(c.expected) {
				t.Errorf("\n got: %s\n want: %s\n", actual, c.expected)
			}
		})
	}
}

func TestFieldElement_Pow(t *testing.T) {
	cases := []struct {
		a   ecc.FieldInterface
		expected *ecc.FieldElement
	}{
		{
			_newFieldElement(3, 13).Pow(big.NewInt(3)),
			_newFieldElement(1, 13),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Ne(c.expected) {
				t.Errorf("\n got: %s\n want: %s\n", actual, c.expected)
			}
		})
	}
}

func TestFieldElement_Div(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		a, err := ecc.NewFieldElementFromInt64(0, 1)
		if err != nil {
			t.Fatal(err)
		}
		if _, err := a.Div(_newFieldElement(0, 3)).Calc(); err == nil {
			t.Error("should fail to division two numbers in different Fields")
		}
	})

	cases := []struct {
		a   ecc.FieldInterface
		expected *ecc.FieldElement
	}{
		{
			_newFieldElement(2, 19).Div(_newFieldElement(7, 19)),
			_newFieldElement(3, 19),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Ne(c.expected) {
				t.Errorf("\n got: %s\n want: %s\n", actual, c.expected)
			}
		})
	}
}
