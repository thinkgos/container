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

// sets.Int32 is a set of int32s, implemented via map[int32]struct{} for minimal memory consumption.
type Int32 map[int32]Empty

// NewInt32 creates a Int32 from a list of values.
func NewInt32(items ...int32) Int32 {
	ss := Int32{}
	ss.Insert(items...)
	return ss
}

// Int32KeySet creates a Int32 from a keys of a map[int32](? extends interface{}).
// If the value passed in is not actually a map, this will panic.
func Int32KeySet(theMap interface{}) Int32 {
	v := reflect.ValueOf(theMap)
	ret := Int32{}

	for _, keyValue := range v.MapKeys() {
		ret.Insert(keyValue.Interface().(int32))
	}
	return ret
}

// Insert adds items to the set.
func (s Int32) Insert(items ...int32) Int32 {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

// Delete removes all items from the set.
func (s Int32) Delete(items ...int32) Int32 {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Contains returns true if and only if item is contained in the set.
func (s Int32) Contains(item int32) bool {
	_, contained := s[item]
	return contained
}

// ContainsAll returns true if and only if all items are contained in the set.
func (s Int32) ContainsAll(items ...int32) bool {
	for _, item := range items {
		if !s.Contains(item) {
			return false
		}
	}
	return true
}

// ContainsAny returns true if any items are contained in the set.
func (s Int32) ContainsAny(items ...int32) bool {
	for _, item := range items {
		if s.Contains(item) {
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
func (s Int32) Difference(s2 Int32) Int32 {
	result := NewInt32()
	for key := range s {
		if !s2.Contains(key) {
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
func (s1 Int32) Union(s2 Int32) Int32 {
	result := NewInt32()
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
func (s1 Int32) Intersection(s2 Int32) Int32 {
	var walk, other Int32
	result := NewInt32()
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for key := range walk {
		if other.Contains(key) {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Int32) IsSuperset(s2 Int32) bool {
	for item := range s2 {
		if !s1.Contains(item) {
			return false
		}
	}
	return true
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 Int32) Equal(s2 Int32) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

type sortableSliceOfInt32 []int32

func (s sortableSliceOfInt32) Len() int           { return len(s) }
func (s sortableSliceOfInt32) Less(i, j int) bool { return lessInt32(s[i], s[j]) }
func (s sortableSliceOfInt32) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// List returns the contents as a sorted int32 slice.
func (s Int32) List() []int32 {
	res := make(sortableSliceOfInt32, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	sort.Sort(res)
	return []int32(res)
}

// UnsortedList returns the slice with contents in random order.
func (s Int32) UnsortedList() []int32 {
	res := make([]int32, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

// PopAny Returns a single element from the set.
func (s Int32) PopAny() (int32, bool) {
	for key := range s {
		s.Delete(key)
		return key, true
	}
	return 0, false
}

// Len returns the size of the set.
func (s Int32) Len() int {
	return len(s)
}

func lessInt32(lhs, rhs int32) bool {
	return lhs < rhs
}
