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

func TestPoint_Add(t *testing.T) {
	a := 5
	b := 7
	cases := []struct{
		a *Point
		b *Point
		expected *Point
	} {
		{
			&Point{a, b, -1, -1, nil},
			&Point{a, b, math.NaN(), math.NaN(), nil},
			&Point{a, b, -1, -1, nil},
		},
		{
			&Point{a, b, math.NaN(), math.NaN(), nil},
			&Point{a, b, -1, 1, nil},
			&Point{a, b, -1, 1, nil},
		},
		{
			&Point{a, b, -1, -1, nil},
			&Point{a, b, -1, 1, nil},
			&Point{a, b, math.NaN(), math.NaN(), nil},
		},
		{
			&Point{a, b, 2, 5, nil},
			&Point{a, b, -1, -1, nil},
			&Point{a, b, 3, -7, nil},
		},
		{
			&Point{a, b, -1, -1, nil},
			&Point{a, b, -1, -1, nil},
			&Point{a, b, 18, 77, nil},
		},
		{
			&Point{a, b, -1, 0, nil},
			&Point{a, b, -1, 0, nil},
			&Point{a, b, math.NaN(), math.NaN(), nil},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Add(c.b).Calc()
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
