// Copyright [2020] [thinkgos]
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package array

// Interface is a type of list, both ArrayList and LinkedList implement this interface.
type Interface interface {
	// Len returns the number of elements in the collection.
	Len() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear initializes or clears all of the elements from this container.
	Clear()

	// Push appends the specified elements to the end of this list.
	Push(vals interface{})
	// PushFront inserts a new element e with value v at the front of list l
	PushFront(v interface{})
	// PushBack inserts a new element e with value v at the back of list l.
	PushBack(v interface{})
	// Add inserts the specified element at the specified position in this list.
	Add(index int, val interface{}) error

	// Poll return the front element value and then remove from list
	Poll() interface{}
	// PollFront return the front element value and then remove from list
	PollFront() interface{}
	// PollBack return the back element value and then remove from list
	PollBack() interface{}
	// Remove removes the element at the specified position in this list.
	// It returns an error if the index is out of range.
	Remove(index int) (interface{}, error)
	// RemoveWithValue removes the first occurence of the specified element from this list, if it is present.
	// It returns false if the target value isn't present, otherwise returns true.
	RemoveWithValue(val interface{}) bool

	// Get returns the element at the specified positon in this list. The index must be in the range of [0, size).
	Get(index int) (interface{}, error)
	// Peek return the front element value
	Peek() interface{}
	// PeekFront return the front element value
	PeekFront() interface{}
	// PeekBack return the back element value
	PeekBack() interface{}

	// Iterator returns an iterator over the elements in this list in proper sequence.
	Iterator(f func(interface{}) bool)
	// ReverseIterator returns an iterator over the elements in this list in reverse sequence as Iterator.
	ReverseIterator(f func(interface{}) bool)

	// Contains returns true if this list contains the specified element.
	Contains(val interface{}) bool
	// Sort sorts the element using default options below. It sorts the elements into ascending sequence according to their natural ordering.
	Sort(reverse ...bool)
	// Values get a copy of all the values in the list
	Values() []interface{}
}
