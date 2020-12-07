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

	"github.com/stretchr/testify/require"
)

func TestByteSet(t *testing.T) {
	s := Byte{}
	s2 := Byte{}
	if len(s) != 0 {
		t.Errorf("Expected len=0: %d", len(s))
	}
	s.Insert(1, 2)
	if len(s) != 2 {
		t.Errorf("Expected len=2: %d", len(s))
	}
	s.Insert(3)
	if s.Contains(4) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Contains(1) {
		t.Errorf("Missing contents: %#v", s)
	}
	s.Delete(1)
	if s.Contains(1) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	s.Insert(1)
	if s.ContainsAll(1, 2, 4) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.ContainsAll(1, 2) {
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
	_, ok := s2.PopAny()
	if !ok {
		t.Errorf("Unexpected status: %#v", ok)
	}
	s2 = NewByte()
	if s2.Len() != 0 {
		t.Errorf("Expected len=0: %d", len(s2))
	}
	v, ok := s2.PopAny()
	if ok {
		t.Errorf("Unexpected status: %#v", ok)
	}
	if v != 0 {
		t.Errorf("Unexpected value: %#v", v)
	}
	// improve cover
	s2 = ByteKeySet(map[byte]interface{}{1: "1", 2: "2", 3: "3"})
	s2.UnsortedList()
}

func TestByteSetDeleteMultiples(t *testing.T) {
	s := Byte{}
	s.Insert(1, 2, 3)
	if len(s) != 3 {
		t.Errorf("Expected len=3: %d", len(s))
	}

	s.Delete(1, 3)
	if len(s) != 1 {
		t.Errorf("Expected len=1: %d", len(s))
	}
	if s.Contains(1) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if s.Contains(3) {
		t.Errorf("Unexpected contents: %#v", s)
	}
	if !s.Contains(2) {
		t.Errorf("Missing contents: %#v", s)
	}
}

func TestNewByteSet(t *testing.T) {
	s := NewByte(1, 2, 3)
	if len(s) != 3 {
		t.Errorf("Expected len=3: %d", len(s))
	}
	if !s.Contains(1) || !s.Contains(2) || !s.Contains(3) {
		t.Errorf("Unexpected contents: %#v", s)
	}
}

func TestByteSetList(t *testing.T) {
	s := NewByte(13, 12, 11, 1)
	if !reflect.DeepEqual(s.List(), []byte{1, 11, 12, 13}) {
		t.Errorf("List gave unexpected result: %#v", s.List())
	}
}

func TestByteSetDifference(t *testing.T) {
	a := NewByte(1, 2, 3)
	b := NewByte(1, 2, 4, 5)
	c := a.Difference(b)
	d := b.Difference(a)
	if len(c) != 1 {
		t.Errorf("Expected len=1: %d", len(c))
	}
	if !c.Contains(3) {
		t.Errorf("Unexpected contents: %#v", c.List())
	}
	if len(d) != 2 {
		t.Errorf("Expected len=2: %d", len(d))
	}
	if !d.Contains(4) || !d.Contains(5) {
		t.Errorf("Unexpected contents: %#v", d.List())
	}
}

func TestByteSetHasAny(t *testing.T) {
	a := NewByte(1, 2, 3)

	if !a.ContainsAny(1, 4) {
		t.Errorf("expected true, got false")
	}

	if a.ContainsAny(10, 4) {
		t.Errorf("expected false, got true")
	}
}

func TestByteSetEquals(t *testing.T) {
	// Simple case (order doesn't matter)
	a := NewByte(1, 2)
	b := NewByte(2, 1)
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// It is a set; duplicates are ignored
	b = NewByte(2, 2, 1)
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	// Edge cases around empty sets / empty strings
	a = NewByte()
	b = NewByte()
	if !a.Equal(b) {
		t.Errorf("Expected to be equal: %v vs %v", a, b)
	}

	b = NewByte(1, 2, 3)
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	b = NewByte(1, 2, 0)
	if a.Equal(b) {
		t.Errorf("Expected to be not-equal: %v vs %v", a, b)
	}

	// Check for equality after mutation
	a = NewByte()
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

func TestByteUnion(t *testing.T) {
	tests := []struct {
		s1       Byte
		s2       Byte
		expected Byte
	}{
		{
			NewByte(1, 2, 3, 4),
			NewByte(3, 4, 5, 6),
			NewByte(1, 2, 3, 4, 5, 6),
		},
		{
			NewByte(1, 2, 3, 4),
			NewByte(),
			NewByte(1, 2, 3, 4),
		},
		{
			NewByte(),
			NewByte(1, 2, 3, 4),
			NewByte(1, 2, 3, 4),
		},
		{
			NewByte(),
			NewByte(),
			NewByte(),
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

func TestByteIntersection(t *testing.T) {
	tests := []struct {
		s1       Byte
		s2       Byte
		expected Byte
	}{
		{
			NewByte(1, 2, 3, 4),
			NewByte(3, 4, 5, 6),
			NewByte(3, 4),
		},
		{
			NewByte(1, 2, 3, 4),
			NewByte(1, 2, 3, 4),
			NewByte(1, 2, 3, 4),
		},
		{
			NewByte(1, 2, 3, 4),
			NewByte(),
			NewByte(),
		},
		{
			NewByte(),
			NewByte(1, 2, 3, 4),
			NewByte(),
		},
		{
			NewByte(),
			NewByte(),
			NewByte(),
		},
	}

	for _, test := range tests {
		intersection := test.s1.Intersection(test.s2)
		if intersection.Len() != test.expected.Len() {
			t.Errorf("Expected intersection.Len()=%d but got %d", test.expected.Len(), intersection.Len())
		}

		if !intersection.Equal(test.expected) {
			t.Errorf("Expected intersection.Equal(expected) but not true.  intersection:%v expected:%v",
				intersection.List(), test.expected.List())
		}
	}
}

func TestByte_Each(t *testing.T) {
	expect := NewByte(1, 2, 3, 4)
	s1 := NewByte(1, 2, 3, 4)
	s1.Each(func(item interface{}) bool {
		require.True(t, expect.Contains(item.(byte)))
		return item.(byte) != 3
	})
}

func TestByte_Clone(t *testing.T) {
	s1 := NewByte(1, 2, 3, 4)
	s2 := s1.Clone()

	require.True(t, s1.Equal(s2))
}
