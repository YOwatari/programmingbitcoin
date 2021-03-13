package ecc_test

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"

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

func TestPoint_Eq(t *testing.T) {
	a := NewExampleInteger(5)
	b := NewExampleInteger(7)

	cases := []struct{
		a map[string]int
		b map[string]int
		expected bool
	} {
		{
			map[string]int{"x": -1, "y": -1},
			map[string]int{"x": -1, "y": -1},
			true,
		},
		{
			map[string]int{"x": -1, "y": -1},
			map[string]int{"x": 18, "y": 77},
			false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			x1 := NewExampleInteger(c.a["x"])
			y1 := NewExampleInteger(c.a["y"])
			actual1, err := ecc.NewPoint(x1, y1, a, b)
			if err != nil {
				t.Fatal(err)
			}

			x2 := NewExampleInteger(c.b["x"])
			y2 := NewExampleInteger(c.b["y"])
			actual2, err := ecc.NewPoint(x2, y2, a, b)
			if err != nil {
				t.Fatal(err)
			}

			if actual1.Eq(actual2) != c.expected {
				t.Errorf("Point.Eq: %#v, %#v, expected: %t", actual1, actual2, c.expected)
			}
		})
	}
}

func TestPoint_Ne(t *testing.T) {
	a := NewExampleInteger(5)
	b := NewExampleInteger(7)

	cases := []struct{
		a map[string]int
		b map[string]int
		expected bool
	} {
		{
			map[string]int{"x": -1, "y": -1},
			map[string]int{"x": -1, "y": -1},
			false,
		},
		{
			map[string]int{"x": -1, "y": -1},
			map[string]int{"x": 18, "y": 77},
			true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			x1 := NewExampleInteger(c.a["x"])
			y1 := NewExampleInteger(c.a["y"])
			actual1, err := ecc.NewPoint(x1, y1, a, b)
			if err != nil {
				t.Fatal(err)
			}

			x2 := NewExampleInteger(c.b["x"])
			y2 := NewExampleInteger(c.b["y"])
			actual2, err := ecc.NewPoint(x2, y2, a, b)
			if err != nil {
				t.Fatal(err)
			}

			if actual1.Ne(actual2) != c.expected {
				t.Errorf("Point.Ne: %#v, %#v, expected: %t", actual1, actual2, c.expected)
			}
		})
	}
}

func TestPoint_Add(t *testing.T) {
	a := NewExampleInteger(5)
	b := NewExampleInteger(7)

	cases := []struct{
		a map[string]interface{}
		b map[string]interface{}
		expected *ecc.Point
	} {
		{
			map[string]interface{}{"x": -1, "y": -1},
			map[string]interface{}{"x": nil, "y": nil},
			&ecc.Point{
				X: NewExampleInteger(-1),
				Y: NewExampleInteger(-1),
				A: a, B: b, Err: nil},
		},
		{
			map[string]interface{}{"x": nil, "y": nil},
			map[string]interface{}{"x": -1, "y": -1},
			&ecc.Point{
				X: NewExampleInteger(-1),
				Y: NewExampleInteger(-1),
				A: a, B: b, Err: nil},
		},
		{
			map[string]interface{}{"x": -1, "y": -1},
			map[string]interface{}{"x": -1, "y": 1},
			&ecc.Point{
				X: nil,
				Y: nil,
				A: a, B: b, Err: nil},
		},
		{
			map[string]interface{}{"x": 2, "y": 5},
			map[string]interface{}{"x": -1, "y": -1},
			&ecc.Point{
				X: NewExampleInteger(3),
				Y: NewExampleInteger(-7),
				A: a, B: b, Err: nil},
		},
		{
			map[string]interface{}{"x": -1, "y": -1},
			map[string]interface{}{"x": -1, "y": -1},
			&ecc.Point{
				X: NewExampleInteger(18),
				Y: NewExampleInteger(77),
				A: a, B: b, Err: nil},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			var x1, y1 ecc.FieldInterface
			if x, ok := c.a["x"].(int); ok {
				x1 = NewExampleInteger(x)
			} else {
				x1 = nil
			}
			if y, ok := c.a["y"].(int); ok {
				y1 = NewExampleInteger(y)
			} else {
				y1 = nil
			}
			p1, err := ecc.NewPoint(x1, y1, a, b)
			if err != nil {
				t.Fatal(err)
			}

			var x2, y2 ecc.FieldInterface
			if x, ok := c.b["x"].(int); ok {
				x2 = NewExampleInteger(x)
			} else {
				x2 = nil
			}
			if y, ok := c.b["y"].(int); ok {
				y2 = NewExampleInteger(y)
			} else {
				y2 = nil
			}
			p2, err := ecc.NewPoint(x2, y2, a, b)
			if err != nil {
				t.Fatal(err)
			}

			actual, err := p1.Add(p2).Calc()
			if err != nil {
				t.Fatal(err)
			}

			if actual.Ne(c.expected) {
				diff := cmp.Diff(actual, c.expected)
				t.Errorf("Point diff: (-got +want)\n%s", diff)
			}
		})
	}
}

type ExampleInteger struct {
	N int
}

func NewExampleInteger(n int) *ExampleInteger {
	return &ExampleInteger{N: n}
}

func (i *ExampleInteger) Eq(other ecc.FieldInterface) bool {
	o := other.(*ExampleInteger)
	return i.N == o.N
}

func (i *ExampleInteger) Ne(other ecc.FieldInterface) bool {
	return !i.Eq(other)
}

func (i *ExampleInteger) Calc() (ecc.FieldInterface, error) {
	return i, nil
}

func (i *ExampleInteger) Copy() ecc.FieldInterface {
	return &ExampleInteger{N: i.N}
}

func (i *ExampleInteger) MulInt(c int) ecc.FieldInterface {
	*i = ExampleInteger{N: i.N * c}
	return i
}

func (i *ExampleInteger) Add(other ecc.FieldInterface) ecc.FieldInterface {
	o := other.(*ExampleInteger)
	*i = ExampleInteger{N: i.N + o.N}
	return i
}

func (i *ExampleInteger) Sub(other ecc.FieldInterface) ecc.FieldInterface {
	o := other.(*ExampleInteger)
	*i = ExampleInteger{N: i.N - o.N}
	return i
}

func (i *ExampleInteger) Mul(other ecc.FieldInterface) ecc.FieldInterface {
	o := other.(*ExampleInteger)
	*i = ExampleInteger{N: i.N * o.N}
	return i
}

func (i *ExampleInteger) Pow(exp int) ecc.FieldInterface {
	*i = ExampleInteger{N: int(math.Pow(float64(i.N), float64(exp)))}
	return i
}

func (i *ExampleInteger) Div(other ecc.FieldInterface) ecc.FieldInterface {
	o := other.(*ExampleInteger)
	*i = ExampleInteger{N: i.N / o.N}
	return i
}
