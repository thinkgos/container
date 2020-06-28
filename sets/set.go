package sets

import (
	"reflect"
)

// Set sets.Set is a set of interface, implemented via map[interface{}]struct{} for minimal memory consumption.
type Set map[interface{}]Empty

// New creates a interface{} from a list of values.
func New(items ...interface{}) Set {
	ss := Set{}
	ss.Insert(items...)
	return ss
}

// SetKey creates a interface{} from a keys of a map[interface{}](? extends interface{}).
// If the value passed in is not actually a map, this will panic.
func SetKey(theMap interface{}) Set {
	v := reflect.ValueOf(theMap)
	ret := Set{}

	for _, keyValue := range v.MapKeys() {
		ret.Insert(keyValue.Interface())
	}
	return ret
}

// Insert adds items to the set.
func (s Set) Insert(items ...interface{}) Set {
	for _, item := range items {
		s[item] = Empty{}
	}
	return s
}

// Delete removes all items from the set.
func (s Set) Delete(items ...interface{}) Set {
	for _, item := range items {
		delete(s, item)
	}
	return s
}

// Contain returns true if and only if item is contained in the set.
func (s Set) Contain(item interface{}) bool {
	_, contained := s[item]
	return contained
}

// ContainAll returns true if and only if all items are contained in the set.
func (s Set) ContainAll(items ...interface{}) bool {
	for _, item := range items {
		if !s.Contain(item) {
			return false
		}
	}
	return true
}

// ContainAny returns true if any items are contained in the set.
func (s Set) ContainAny(items ...interface{}) bool {
	for _, item := range items {
		if s.Contain(item) {
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
func (s Set) Difference(s2 Set) Set {
	result := New()
	for key := range s {
		if !s2.Contain(key) {
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
func (s1 Set) Union(s2 Set) Set {
	result := New()
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
func (s1 Set) Intersection(s2 Set) Set {
	var walk, other Set
	result := New()
	if s1.Len() < s2.Len() {
		walk = s1
		other = s2
	} else {
		walk = s2
		other = s1
	}
	for key := range walk {
		if other.Contain(key) {
			result.Insert(key)
		}
	}
	return result
}

// IsSuperset returns true if and only if s1 is a superset of s2.
func (s1 Set) IsSuperset(s2 Set) bool {
	for item := range s2 {
		if !s1.Contain(item) {
			return false
		}
	}
	return true
}

// List returns the contents as a sorted slice.
func (s Set) List() []interface{} {
	// TODO: implement list
	panic("not implement")
	return nil
}

// UnsortedList returns the slice with contents in random order.
func (s Set) UnsortedList() []interface{} {
	res := make([]interface{}, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

// Equal returns true if and only if s1 is equal (as a set) to s2.
// Two sets are equal if their membership is identical.
// (In practice, this means same elements, order doesn't matter)
func (s1 Set) Equal(s2 Set) bool {
	return len(s1) == len(s2) && s1.IsSuperset(s2)
}

// PopAny Returns a single element from the set.
func (s Set) PopAny() (interface{}, bool) {
	for key := range s {
		s.Delete(key)
		return key, true
	}
	var zeroValue interface{}
	return zeroValue, false
}

// Len returns the size of the set.
func (s Set) Len() int {
	return len(s)
}
