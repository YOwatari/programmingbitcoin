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
	a, err := NewFieldElement(7, 13)
	if err != nil {
		t.Fatal(err)
	}
	b, err := NewFieldElement(6, 13)
	if err != nil {
		t.Fatal(err)
	}

	if a.Eq(b) != false {
		t.Errorf("%#v should not be equal to %#v", a, b)
	}

	if a.Eq(a) != true {
		t.Errorf("%#v should be equal to %#v", a, a)
	}
}
