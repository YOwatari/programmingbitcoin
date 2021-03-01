package ecc

import (
	"fmt"
	"math"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPoint(t *testing.T) {
	t.Run("Succeeds", func(t *testing.T) {
		actual, err := NewPoint(-1, -1, 5, 7)
		if err != nil {
			t.Fatal(err)
		}
		expected, _ := NewPoint(-1, -1, 5, 7)

		if diff := cmp.Diff(actual, expected); diff != "" {
			t.Errorf("Point diff: (-got +want)\n%s", diff)
		}
	})

	t.Run("Fails", func(t *testing.T) {
		actual, err := NewPoint(-1, -2, 5, 7)
		if err == nil || actual != nil {
			t.Error("should be failed")
		}
	})
}

func TestPoint_Eq(t *testing.T) {
	cases := []struct{
		a *Point
		b *Point
		expected bool
	} {
		{
			&Point{5, 7, -1, -1, nil},
			&Point{5, 7, -1, -1, nil},
			true,
		},
		{
			&Point{5, 7, -1, -1, nil},
			&Point{5, 7, 18, 57, nil},
			false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if c.a.Eq(c.b) != c.expected {
				t.Errorf("Point.Eq: %#v, %#v, expected: %t", c.a, c.b, c.expected)
			}
		})
	}
}

func TestPoint_Ne(t *testing.T) {
	cases := []struct{
		a *Point
		b *Point
		expected bool
	} {
		{
			&Point{5, 7, -1, -1, nil},
			&Point{5, 7, -1, -1, nil},
			false,
		},
		{
			&Point{5, 7, -1, -1, nil},
			&Point{5, 7, 18, 57, nil},
			true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if c.a.Ne(c.b) != c.expected {
				t.Errorf("Point.Ne: %#v, %#v, expected: %t", c.a, c.b, c.expected)
			}
		})
	}
}

func TestPoint_Add_inf(t *testing.T) {
	p1, _ := NewPoint(-1, -1, 5, 7)
	p2, _ := NewPoint(-1, 1, 5, 7)
	inf, _ := NewPoint(math.NaN(), math.NaN(), 5, 7)

	a1, err := p1.Add(inf).Calc()
	if err != nil {
		t.Fatal(err)
	}
	if a1.Ne(p1) {
		t.Errorf("p1 + inf: got: %v, want: %v", a1, p1)
	}

	a2, err := inf.Add(p2).Calc()
	if err != nil {
		t.Fatal(err)
	}
	if a2.Ne(p2) {
		t.Errorf("inf + p2: got: %v, want: %v", a2, p2)
	}

	a3, err := p1.Add(p2).Calc()
	if err != nil {
		t.Fatal(err)
	}
	if a3.Ne(inf) {
		t.Errorf("p1 + p2: got: %v, want: %v", a3, inf)
	}
}
