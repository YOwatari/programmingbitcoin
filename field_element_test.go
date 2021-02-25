package main

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewFieldElement_Succeeds(t *testing.T) {
	actual, err := NewFieldElement(1, 10)
	if err != nil {
		t.Error(err)
	}
	expected := &FieldElement{
		Num: 1,
		Prime: 10,
	}

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Errorf("FieldElement diff: (-got +want)\n%s", diff)
	}
}

func TestNewFieldElement_Fails(t *testing.T) {
	cases := []struct{
		num int
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
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			got, err := NewFieldElement(c.num, c.prime)
			if err == nil || got != nil {
				t.Error("should fail")
			}
		})
	}
}

func TestFieldElement_Eq(t *testing.T) {
	cases := []struct{
		a *FieldElement
		b *FieldElement
		expected bool
	} {
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
			&FieldElement{7, 12},
			false,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if c.a.Eq(c.b) != c.expected {
				t.Errorf("FieldElement.Eq: %#v, %#v expected: %t", c.a, c.b, c.expected)
			}
		})
	}
}

func TestFieldElement_Ne(t *testing.T) {
	cases := []struct{
		a *FieldElement
		b *FieldElement
		expected bool
	} {
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
			&FieldElement{7, 12},
			true,
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			if c.a.Ne(c.b) != c.expected {
				t.Errorf("FieldElement.Ne: %#v, %#v expected: %t", c.a, c.b, c.expected)
			}
		})
	}
}
