package ecc_test

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/YOwatari/programmingbitcoin/ecc"
)

func TestNewPoint(t *testing.T) {
	t.Run("Succeeds", func(t *testing.T) {
		_, err := ecc.NewPoint(
			NewExampleInteger(-1),
			NewExampleInteger(-1),
			NewExampleInteger(5),
			NewExampleInteger(7),
		)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Fails", func(t *testing.T) {
		actual, err := ecc.NewPoint(
			NewExampleInteger(-1),
			NewExampleInteger(-2),
			NewExampleInteger(5),
			NewExampleInteger(7),
		)
		if err == nil || actual != nil {
			t.Error("should be failed")
		}
	})
}

func TestNewPoint_FieldElement(t *testing.T) {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)

	t.Run("Succeeds", func(t *testing.T) {
		for _, v := range []struct {
			x int64
			y int64
		} {
			{192, 105},
			{17, 56},
			{1, 193},
		} {
			x, err := ecc.NewFieldElementFromInt64(v.x, prime)
			if err != nil {
				t.Fatal(err)
			}
			y, err := ecc.NewFieldElementFromInt64(v.y, prime)
			if err != nil {
				t.Fatal(err)
			}

			_, err = ecc.NewPoint(x, y, a, b)
			if err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Fails", func(t *testing.T) {
		for _, v := range []struct {
			x int64
			y int64
		} {
			{200, 119},
			{42, 99},
		} {
			x, err := ecc.NewFieldElementFromInt64(v.x, prime)
			if err != nil {
				t.Fatal(err)
			}
			y, err := ecc.NewFieldElementFromInt64(v.y, prime)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := ecc.NewPoint(x, y, a, b)
			if err == nil || actual != nil {
				t.Error(err)
			}
		}
	})
}

func TestPoint_Eq(t *testing.T) {
	a := int64(5)
	b := int64(7)

	cases := []struct {
		p1       *ecc.Point
		p2       *ecc.Point
		expected bool
	}{
		{
			_NewExampleIntegerPoint(-1, -1, a, b),
			_NewExampleIntegerPoint(-1, -1, a, b),
			true,
		},
		{
			_NewExampleIntegerPoint(-1, -1, a, b),
			_NewExampleIntegerPoint(18, 77, a, b),
			false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if c.p1.Eq(c.p2) != c.expected {
				t.Errorf("Point.Eq: %#v, %#v, expected: %t", c.p1, c.p2, c.expected)
			}
		})
	}
}

func TestPoint_Ne(t *testing.T) {
	a := int64(5)
	b := int64(7)

	cases := []struct {
		p1       *ecc.Point
		p2       *ecc.Point
		expected bool
	}{
		{
			_NewExampleIntegerPoint(-1, -1, a, b),
			_NewExampleIntegerPoint(-1, -1, a, b),
			false,
		},
		{
			_NewExampleIntegerPoint(-1, -1, a, b),
			_NewExampleIntegerPoint(18, 77, a, b),
			true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if c.p1.Ne(c.p2) != c.expected {
				t.Errorf("Point.Ne: %#v, %#v, expected: %t", c.p1, c.p2, c.expected)
			}
		})
	}
}

func TestPoint_Add(t *testing.T) {
	a := int64(5)
	b := int64(7)

	cases := []struct {
		p1       *ecc.Point
		p2       *ecc.Point
		expected *ecc.Point
	}{
		{
			_NewExampleIntegerPoint(-1, -1, a, b),
			&ecc.Point{X: nil, Y: nil, A: NewExampleInteger(a), B: NewExampleInteger(b)},
			_NewExampleIntegerPoint(-1, -1, a, b),
		},
		{
			&ecc.Point{X: nil, Y: nil, A: NewExampleInteger(a), B: NewExampleInteger(b)},
			_NewExampleIntegerPoint(-1, -1, a, b),
			_NewExampleIntegerPoint(-1, -1, a, b),
		},
		{
			_NewExampleIntegerPoint(-1, -1, a, b),
			_NewExampleIntegerPoint(-1, 1, a, b),
			&ecc.Point{X: nil, Y: nil, A: NewExampleInteger(a), B: NewExampleInteger(b)},
		},
		{
			_NewExampleIntegerPoint(2, 5, a, b),
			_NewExampleIntegerPoint(-1, -1, a, b),
			_NewExampleIntegerPoint(3, -7, a, b),
		},
		{
			_NewExampleIntegerPoint(-1, -1, a, b),
			_NewExampleIntegerPoint(-1, -1, a, b),
			_NewExampleIntegerPoint(18, 77, a, b),
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.p1.Add(c.p1, c.p2).Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Ne(c.expected) {
				t.Errorf("\n got: %s\nwant: %s", actual, c.expected)
			}
		})
	}
}

func TestPoint_Add_FieldElement(t *testing.T) {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)

	type point struct {
		x int64
		y int64
	}
	cases := []struct{
		a point
		b point
		expected *ecc.Point
	} {
		{
			point{192, 105},
			point{17, 56},
			&ecc.Point{
				X: _newFieldElement(170, prime),
				Y: _newFieldElement(142, prime),
				A: a, B: b, Err: nil,
			},
		},
		{
			point{170, 142},
			point{60, 139},
			&ecc.Point{
				X: _newFieldElement(220, prime),
				Y: _newFieldElement(181, prime),
				A: a, B: b, Err: nil,
			},
		},
		{
			point{47, 71},
			point{17, 56},
			&ecc.Point{
				X: _newFieldElement(215, prime),
				Y: _newFieldElement(68, prime),
				A: a, B: b, Err: nil,
			},
		},
		{
			point{143, 98},
			point{76, 66},
			&ecc.Point{
				X: _newFieldElement(47, prime),
				Y: _newFieldElement(71, prime),
				A: a, B: b, Err: nil,
			},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			x1, err := ecc.NewFieldElementFromInt64(c.a.x, prime)
			if err != nil {
				t.Fatal(err)
			}
			y1, err := ecc.NewFieldElementFromInt64(c.a.y, prime)
			if err != nil {
				t.Fatal(err)
			}
			p1, err := ecc.NewPoint(x1, y1, a, b)
			if err != nil {
				t.Fatal(err)
			}

			x2, err := ecc.NewFieldElementFromInt64(c.b.x, prime)
			if err != nil {
				t.Fatal(err)
			}
			y2, err := ecc.NewFieldElementFromInt64(c.b.y, prime)
			if err != nil {
				t.Fatal(err)
			}
			p2, err := ecc.NewPoint(x2, y2, a, b)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := p1.Add(p1, p2).Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Ne(c.expected) {
				t.Errorf("\n got: %s\n want: %s\n", actual, c.expected)
			}
		})
	}
}

func TestPoint_RMul(t *testing.T) {
	prime := int64(223)
	a, _ := ecc.NewFieldElementFromInt64(0, prime)
	b, _ := ecc.NewFieldElementFromInt64(7, prime)
	x, _ := ecc.NewFieldElementFromInt64(15, prime)
	y, _ := ecc.NewFieldElementFromInt64(86, prime)
	expected := &ecc.Point{X: nil, Y: nil, A: a, B: b}

	p, err := ecc.NewPoint(x, y, a, b)
	if err != nil {
		t.Fatal(err)
	}

	actual, err := p.RMul(p, big.NewInt(7)).Calc()
	if err != nil {
		t.Fatal(err)
	}

	if actual.Ne(expected) {
		t.Errorf("\n got: %s\n want: %s\n", actual, expected)
	}
}

func BenchmarkPoint_RMul(b *testing.B) {
	prime := int64(223)
	A, _ := ecc.NewFieldElementFromInt64(0, prime)
	B, _ := ecc.NewFieldElementFromInt64(7, prime)
	x, _ := ecc.NewFieldElementFromInt64(15, prime)
	y, _ := ecc.NewFieldElementFromInt64(86, prime)
	p, err := ecc.NewPoint(x, y, A, B)
	if err != nil {
		b.Fatal(err)
	}

	n := new(big.Int).Lsh(big.NewInt(1), 40)
	b.ResetTimer()
	if _, err := p.RMul(p, n).Calc(); err != nil {
		b.Fatal(err)
	}
}

type ExampleInteger struct {
	Num int64
}

func NewExampleInteger(n int64) *ExampleInteger {
	return &ExampleInteger{Num: n}
}

func (i *ExampleInteger) Eq(other ecc.FieldInterface) bool {
	o := other.(*ExampleInteger)
	return i.Num == o.Num
}

func (i *ExampleInteger) Ne(other ecc.FieldInterface) bool {
	return !i.Eq(other)
}

func (i *ExampleInteger) Calc() (ecc.FieldInterface, error) {
	return i, nil
}

func (i *ExampleInteger) Copy() ecc.FieldInterface {
	return &ExampleInteger{Num: i.Num}
}

func (i *ExampleInteger) Add(a, b ecc.FieldInterface) ecc.FieldInterface {
	x, y := a.(*ExampleInteger), b.(*ExampleInteger)
	*i = ExampleInteger{Num: x.Num + y.Num}
	return i
}

func (i *ExampleInteger) Sub(a, b ecc.FieldInterface) ecc.FieldInterface {
	x, y := a.(*ExampleInteger), b.(*ExampleInteger)
	*i = ExampleInteger{Num: x.Num - y.Num}
	return i
}

func (i *ExampleInteger) Mul(a, b ecc.FieldInterface) ecc.FieldInterface {
	x, y := a.(*ExampleInteger), b.(*ExampleInteger)
	*i = ExampleInteger{Num: x.Num * y.Num}
	return i
}

func (i *ExampleInteger) Div(a, b ecc.FieldInterface) ecc.FieldInterface {
	x, y := a.(*ExampleInteger), b.(*ExampleInteger)
	*i = ExampleInteger{Num: x.Num / y.Num}
	return i
}

func (i *ExampleInteger) Pow(n ecc.FieldInterface, exp *big.Int) ecc.FieldInterface {
	x := n.(*ExampleInteger)
	result := new(big.Int)
	result.Exp(big.NewInt(x.Num), exp, nil)
	*i = ExampleInteger{Num: result.Int64()}
	return i
}

func (i *ExampleInteger) RMul(n ecc.FieldInterface, c *big.Int) ecc.FieldInterface {
	x := n.(*ExampleInteger)
	*i = ExampleInteger{Num: x.Num * c.Int64()}
	return i
}

func _NewExampleIntegerPoint(x int64, y int64, a int64, b int64) *ecc.Point {
	p, err := ecc.NewPoint(NewExampleInteger(x), NewExampleInteger(y), NewExampleInteger(a), NewExampleInteger(b))
	if err != nil {
		panic(err)
	}
	return p
}
