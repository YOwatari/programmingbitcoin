package ecc

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewPoint(t *testing.T) {
	t.Run("Succeeds", func(t *testing.T) {
		actual, err := NewPoint(-1, -1, 5, 7)
		if err != nil {
			t.Fatal(err)
		}
		expected := &Point{5, 7, -1, -1}

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
			&Point{-1, -1, 5, 7},
			&Point{-1, -1, 5, 7},
			true,
		},
		{
			&Point{-1, -1, 5, 7},
			&Point{18, 57, 5, 7},
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
			&Point{-1, -1, 5, 7},
			&Point{-1, -1, 5, 7},
			false,
		},
		{
			&Point{-1, -1, 5, 7},
			&Point{18, 57, 5, 7},
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
