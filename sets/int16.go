/*
Copyright The Kubernetes Authors.

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

// Code generated by set-gen. DO NOT EDIT.

package sets

import (
	"reflect"
	"sort"
)

// sets.Int16 is a set of int16s, implemented via map[int16]struct{} for minimal memory consumption.
type Int16 map[int16]Empty

// NewInt16 creates a Int16 from a list of values.
func NewInt16(items ...int16) Int16 {
	ss := Int16{}
	ss.Insert(items...)
	return ss
}

// Int16KeySet creates a Int16 from a keys of a map[int16](? extends interface{}).
// If the value passed in is not actually a map, this will panic.
func Int16KeySet(theMap interface{}) Int16 {
	v := reflect.ValueOf(theMap)
	ret := Int16{}

	for _, keyValue := range v.MapKeys() {
		ret.Insert(keyValue.Interface().(int16))
	}
	return ret
}

// Insert adds items to the set.
func (s Int16) Insert(items ...int16) Int16 {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

// Delete removes all items from the set.
func (s Int16) Delete(items ...int16) Int16 {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Has returns true if and only if item is contained in the set.
func (s Int16) Has(item int16) bool {
	_, contained := s[item]
	return contained
}

// HasAll returns true if and only if all items are contained in the set.
func (s Int16) HasAll(items ...int16) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any items are contained in the set.
func (s Int16) HasAny(items ...int16) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

// Difference returns a set of objects that are not in s2
// For example:
// s1 = {a1, a2, a3}
// s2 = {a1, a2, a4, a5}
// s1.Difference(s2) = {a3}
// s2.Difference(s1) = {a4, a5}
func (s Int16) Difference(s2 Int16) Int16 {
	result := NewInt16()
	for key := range s {
		if !s2.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// Union returns a new set which includes items in either s1 or s2.
// For example:
// s1 = {a1, a2}
// s2 = {a3, a4}
// s1.Union(s2) = {a1, a2, a3, a4}
// s2.Union(s1) = {a1, a2, a3, a4}
func (s1 Int16) Union(s2 Int16) Int16 {
	result := NewInt16()
	for key := range s1 {
		result.Insert(key)
	}
	for key := range s2 {
		result.Insert(key)
	}
	return result
}

// Intersection returns a new set which includes the item in BOTH s1 and s2
// For example:
// s1 = {a1, a2}
// s2 = {a2, a3}
// s1.Intersection(s2) = {a2}
func (s1 Int16) Intersection(s2 Int16) Int16 {
	var walk, other Int16
	result := NewInt16()
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for key := range walk {
		if other.Has(key) {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Int16) IsSuperset(s2 Int16) bool {
	for item := range s2 {
		if !s1.Has(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 Int16) Equal(s2 Int16) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

type sortableSliceOfInt16 []int16

func (s sortableSliceOfInt16) Len() int           { return len(s) }
func (s sortableSliceOfInt16) Less(i, j int) bool { return lessInt16(s[i], s[j]) }
func (s sortableSliceOfInt16) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// List returns the contents as a sorted int16 slice.
func (s Int16) List() []int16 {
	res := make(sortableSliceOfInt16, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Sort(res)
	return []int16(res)
}

// UnsortedList returns the slice with contents in random order.
func (s Int16) UnsortedList() []int16 {
	res := make([]int16, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

// Returns a single element from the set.
func (s Int16) PopAny() (int16, bool) {
	for key := range s {
		s.Delete(key)
		return key, true
	}
	var zeroValue int16
	return zeroValue, false
}

// Len returns the size of the set.
func (s Int16) Len() int {
	return len(s)
}

func lessInt16(lhs, rhs int16) bool {
	return lhs < rhs
}
