/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sets

import (
	"testing"
)

func TestSet(t *testing.T) {
	s := New()
	s2 := New()
	if len(s.m) != 0 {
		t.Errorf("Expected len=0: %d", len(s.m))
	}
	s.Insert("a", "b")
	if len(s.m) != 2 {
		t.Errorf("Expected len=2: %d", len(s.m))
	}
	s.Insert("c")
	if s.Contain("d") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Contain("a") {
		t.Errorf("Missing contents: %#v", s)
	}
	s.Delete("a")
	if s.Contain("a") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s.Insert("a")
	if s.ContainAll("a", "b", "d") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.ContainAll("a", "b") {
		t.Errorf("Missing contents: %#v", s)
	}
	s2.Insert("a", "b", "d")
	if s.IsSuperset(s2) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s2.Delete("d")
	if !s.IsSuperset(s2) {
		t.Errorf("Missing contents: %#v", s)
	}
}

func TestSetDeleteMultiples(t *testing.T) {
	s := New()
	s.Insert("a", "b", "c")
	if len(s.m) != 3 {
		t.Errorf("Expected len=3: %d", len(s.m))
	}

	s.Delete("a", "c")
	if len(s.m) != 1 {
		t.Errorf("Expected len=1: %d", len(s.m))
	}
	if s.Contain("a") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if s.Contain("c") {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Contain("b") {
		t.Errorf("Missing contents: %#v", s)
	}

}

func TestNewSet(t *testing.T) {
	s := New(WithItems("a", "b", "c"))
	if len(s.m) != 3 {
		t.Errorf("Expected len=3: %d", len(s.m))
	}
	if !s.Contain("a") || !s.Contain("b") || !s.Contain("c") {
		t.Errorf("Unexpected contents: %#v", s)
	}
}

func TestSetDifference(t *testing.T) {
	a := New(WithItems("1", "2", "3"))
	b := New(WithItems("1", "2", "4", "5"))
	c := a.Difference(b)
	d := b.Difference(a)
	if len(c.m) != 1 {
		t.Errorf("Expected len=1: %d", len(c.m))
	}
	if !c.Contain("3") {
		t.Errorf("Unexpected contents: %#v", c.List())
	}
	if len(d.m) != 2 {
		t.Errorf("Expected len=2: %d", len(d.m))
	}
	if !d.Contain("4") || !d.Contain("5") {
		t.Errorf("Unexpected contents: %#v", d.List())
	}
}

func TestSetHasAny(t *testing.T) {
	a := New(WithItems("1", "2", "3"))

	if !a.ContainAny("1", "4") {
		t.Errorf("expected true, got false")
	}

	if a.ContainAny("0", "4") {
		t.Errorf("expected false, got true")
	}
}

func TestSetEquals(t *testing.T) {
	// Simple case (order doesn't matter)
	a := New(WithItems("1", "2"))
	b := New(WithItems("2", "1"))
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// It is a set; duplicates are ignored
	b = New(WithItems("2", "2", "1"))
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// Edge cases around empty sets / empty strings
	a = New()
	b = New()
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	b = New(WithItems("1", "2", "3"))
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	b = New(WithItems("1", "2", ""))
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	// Check for equality after mutation
	a = New()
	a.Insert("1")
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	a.Insert("2")
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	a.Insert("")
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	a.Delete("")
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}
}

func TestUnion(t *testing.T) {
	tests := []struct {
		s1       Set
		s2       Set
		expected Set
	}{
		{
			New(WithItems("1", "2", "3", "4")),
			New(WithItems("3", "4", "5", "6")),
			New(WithItems("1", "2", "3", "4", "5", "6")),
		},
		{
			New(WithItems("1", "2", "3", "4")),
			New(),
			New(WithItems("1", "2", "3", "4")),
		},
		{
			New(),
			New(WithItems("1", "2", "3", "4")),
			New(WithItems("1", "2", "3", "4")),
		},
		{
			New(),
			New(),
			New(),
		},
	}

	for _, test := range tests {
		union := test.s1.Union(test.s2)
		if union.Len() != test.expected.Len() {
			t.Errorf("Expected union.Len()=%d but got %d", test.expected.Len(), union.Len())
		}

		if !union.Equal(test.expected) {
			t.Errorf("Expected union.Equal(expected) but not true.  union:%v expected:%v", union.List(), test.expected.List())
		}
	}
}

func TestIntersection(t *testing.T) {
	tests := []struct {
		s1       Set
		s2       Set
		expected Set
	}{
		{
			New(WithItems("1", "2", "3", "4")),
			New(WithItems("3", "4", "5", "6")),
			New(WithItems("3", "4")),
		},
		{
			New(WithItems("1", "2", "3", "4")),
			New(WithItems("1", "2", "3", "4")),
			New(WithItems("1", "2", "3", "4")),
		},
		{
			New(WithItems("1", "2", "3", "4")),
			New(),
			New(),
		},
		{
			New(),
			New(WithItems("1", "2", "3", "4")),
			New(),
		},
		{
			New(),
			New(),
			New(),
		},
	}

	for _, test := range tests {
		intersection := test.s1.Intersection(test.s2)
		if intersection.Len() != test.expected.Len() {
			t.Errorf("Expected intersection.Len()=%d but got %d", test.expected.Len(), intersection.Len())
		}

		if !intersection.Equal(test.expected) {
			t.Errorf("Expected intersection.Equal(expected) but not true.  intersection:%v expected:%v", intersection.List(), test.expected.List())
		}
	}
}