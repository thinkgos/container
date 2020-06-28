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
	"reflect"
	"testing"
)

func TestInt64Set(t *testing.T) {
	s := Int64{}
	s2 := Int64{}
	if len(s) != 0 {
		t.Errorf("Expected len=0: %d", len(s))
	}
	s.Insert(1, 2)
	if len(s) != 2 {
		t.Errorf("Expected len=2: %d", len(s))
	}
	s.Insert(3)
	if s.Has(4) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Has(1) {
		t.Errorf("Missing contents: %#v", s)
	}
	s.Delete(1)
	if s.Has(1) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s.Insert(1)
	if s.HasAll(1, 2, 4) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.HasAll(1, 2) {
		t.Errorf("Missing contents: %#v", s)
	}
	s2.Insert(1, 2, 4)
	if s.IsSuperset(s2) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s2.Delete(4)
	if !s.IsSuperset(s2) {
		t.Errorf("Missing contents: %#v", s)
	}
}

func TestInt64SetDeleteMultiples(t *testing.T) {
	s := Int64{}
	s.Insert(1, 2, 3)
	if len(s) != 3 {
		t.Errorf("Expected len=3: %d", len(s))
	}

	s.Delete(1, 3)
	if len(s) != 1 {
		t.Errorf("Expected len=1: %d", len(s))
	}
	if s.Has(1) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if s.Has(3) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Has(2) {
		t.Errorf("Missing contents: %#v", s)
	}

}

func TestNewInt64Set(t *testing.T) {
	s := NewInt64(1, 2, 3)
	if len(s) != 3 {
		t.Errorf("Expected len=3: %d", len(s))
	}
	if !s.Has(1) || !s.Has(2) || !s.Has(3) {
		t.Errorf("Unexpected contents: %#v", s)
	}
}

func TestInt64SetList(t *testing.T) {
	s := NewInt64(13, 12, 11, 1)
	if !reflect.DeepEqual(s.List(), []int64{1, 11, 12, 13}) {
		t.Errorf("List gave unexpected result: %#v", s.List())
	}
}

func TestInt64SetDifference(t *testing.T) {
	a := NewInt64(1, 2, 3)
	b := NewInt64(1, 2, 4, 5)
	c := a.Difference(b)
	d := b.Difference(a)
	if len(c) != 1 {
		t.Errorf("Expected len=1: %d", len(c))
	}
	if !c.Has(3) {
		t.Errorf("Unexpected contents: %#v", c.List())
	}
	if len(d) != 2 {
		t.Errorf("Expected len=2: %d", len(d))
	}
	if !d.Has(4) || !d.Has(5) {
		t.Errorf("Unexpected contents: %#v", d.List())
	}
}

func TestInt64SetHasAny(t *testing.T) {
	a := NewInt64(1, 2, 3)

	if !a.HasAny(1, 4) {
		t.Errorf("expected true, got false")
	}

	if a.HasAny(10, 4) {
		t.Errorf("expected false, got true")
	}
}

func TestInt64SetEquals(t *testing.T) {
	// Simple case (order doesn't matter)
	a := NewInt64(1, 2)
	b := NewInt64(2, 1)
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// It is a set; duplicates are ignored
	b = NewInt64(2, 2, 1)
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// Edge cases around empty sets / empty strings
	a = NewInt64()
	b = NewInt64()
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	b = NewInt64(1, 2, 3)
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	b = NewInt64(1, 2, 0)
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	// Check for equality after mutation
	a = NewInt64()
	a.Insert(1)
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	a.Insert(2)
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	a.Insert(0)
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	a.Delete(0)
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}
}

func TestInt64Union(t *testing.T) {
	tests := []struct {
		s1       Int64
		s2       Int64
		expected Int64
	}{
		{
			NewInt64(1, 2, 3, 4),
			NewInt64(3, 4, 5, 6),
			NewInt64(1, 2, 3, 4, 5, 6),
		},
		{
			NewInt64(1, 2, 3, 4),
			NewInt64(),
			NewInt64(1, 2, 3, 4),
		},
		{
			NewInt64(),
			NewInt64(1, 2, 3, 4),
			NewInt64(1, 2, 3, 4),
		},
		{
			NewInt64(),
			NewInt64(),
			NewInt64(),
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

func TestInt64Intersection(t *testing.T) {
	tests := []struct {
		s1       Int64
		s2       Int64
		expected Int64
	}{
		{
			NewInt64(1, 2, 3, 4),
			NewInt64(3, 4, 5, 6),
			NewInt64(3, 4),
		},
		{
			NewInt64(1, 2, 3, 4),
			NewInt64(1, 2, 3, 4),
			NewInt64(1, 2, 3, 4),
		},
		{
			NewInt64(1, 2, 3, 4),
			NewInt64(),
			NewInt64(),
		},
		{
			NewInt64(),
			NewInt64(1, 2, 3, 4),
			NewInt64(),
		},
		{
			NewInt64(),
			NewInt64(),
			NewInt64(),
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
