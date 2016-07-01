// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func TestExample_one(t *testing.T) {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	expected := "{1 9 144}"
	//fmt.Println(x.String()) // "{1 9 144}"
	if x.String() != expected {
		t.Errorf("intset error! Got %s, expected %s", x.String(), expected)
	}

	y.Add(9)
	y.Add(42)
	//fmt.Println(y.String()) // "{9 42}"
	expected = "{9 42}"
	if y.String() != expected {
		t.Errorf("intset error! Got %s, expected %s", y.String(), expected)
	}

	x.UnionWith(&y)
	//fmt.Println(x.String()) // "{1 9 42 144}"
	expected = "{1 9 42 144}"
	if x.String() != expected {
		t.Errorf("intset error! Got %s, expected %s", x.String(), expected)
	}

	//fmt.Println(x.Has(9), x.Has(123)) // "true false"
	if x.Has(9) != true {
		t.Error("intset error! Expected true")
	}
	if x.Has(123) != false {
		t.Error("intset error! Expected false")
	}

	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func TestExample_two(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	expectedVal := "{1 9 42 144}"
	expectedPtr := "{[4398046511618 0 65536]}"
	if fmt.Sprintf("%s", &x) != expectedVal {
		t.Errorf("intset error! Got %s, expected %s", &x, expectedVal)
	}
	if x.String() != expectedVal {
		t.Errorf("intset error! Got %s, expected %s", x.String(), expectedVal)
	}
	if fmt.Sprintf("%v", x) != expectedPtr {
		t.Errorf("intset error! Got %s, expected %s", fmt.Sprintf("%v", x), expectedPtr)
	}
	//!+note
	//fmt.Println(&x)         // "{1 9 42 144}"
	//fmt.Println(x.String()) // "{1 9 42 144}"
	//fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
func TestLen(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	if x.Len() != 4 {
		t.Errorf("intset error: len! Got %d, expected %d", x.Len(), 4)
	}
}

func TestClear(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Clear()

	if x.Len() != 0 {
		t.Errorf("intset error: Clear! Got %d, expected %d", x.Len(), 0)
	}
}

func TestRemove(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(7)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	x.Remove(9)

	if x.Len() != 4 {
		t.Errorf("intset error: Remove! Got %d, expected %d", x.Len(), 4)
	}
}
func TestCopy(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(7)

	y := x.Copy()
	y.Add(34)
	if y.Len() == x.Len() {
		t.Errorf("intset error: Copy! Got %d, expected %d", y.Len(), x.Len())
	}
}
func TestElems(t *testing.T) {
	var x IntSet
	x.Add(1)
	x.Add(7)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	len := 0
	for _ = range x.Elems() {
		len++
	}
	if len != x.Len() {
		t.Errorf("intset error: Elems! Got %d elements, expected %d", len, x.Len())
	}
}
