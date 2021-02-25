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
			&FieldElement{7, 11},
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
			&FieldElement{7, 11},
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

func TestFieldElement_Add(t *testing.T) {
	t.Run("Fails", func(t *testing.T) {
		a := &FieldElement{0, 1}
		b := &FieldElement{0, 3}
		actual, err := a.Add(b)
		if err == nil || actual != nil {
			t.Errorf("should fail to add two numbers in different Fields")
		}
	})

	cases := []struct{
		a *FieldElement
		b *FieldElement
		expected *FieldElement
	} {
		{
			&FieldElement{44, 57},
			&FieldElement{33, 57},
			&FieldElement{20, 57},
		},
		{
			&FieldElement{17, 57},
			&FieldElement{42, 57},
			&FieldElement{2, 57},
		},
		{
			&FieldElement{2, 57},
			&FieldElement{49, 57},
			&FieldElement{51, 57},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Add(c.b)
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
		a := &FieldElement{0, 1}
		b := &FieldElement{0, 3}
		actual, err := a.Sub(b)
		if err == nil || actual != nil {
			t.Errorf("should fail to sub two numbers in different Fields")
		}
	})

	cases := []struct{
		a *FieldElement
		b *FieldElement
		expected *FieldElement
	} {
		{
			&FieldElement{9, 57},
			&FieldElement{29, 57},
			&FieldElement{37, 57},
		},
		{
			&FieldElement{52, 57},
			&FieldElement{30, 57},
			&FieldElement{22, 57},
		},
		{
			&FieldElement{22, 57},
			&FieldElement{38, 57},
			&FieldElement{41, 57},
		},
	}

	for i, c := range cases {
		t.Run(fmt.Sprintf("%d", i), func(t *testing.T) {
			actual, err := c.a.Sub(c.b)
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
