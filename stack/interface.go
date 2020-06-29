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

// Package stack implements a Stack, which orders elements in a LIFO (last-in-first-out) manner.
package stack

// Interface is a Stack, which is LIFO (last-in-first-out).
type Interface interface {
	// Len returns the number of elements in the collection.
	Len() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear initializes or clears all of the elements from this container.
	Clear()
	// Push pushes an element into this Stack.
	Push(interface{})
	// Pop pops the element on the top of this Stack.
	Pop() interface{}
	// Peek retrieves, but does not remove, the element on the top of this Stack, or return nil if this Stack is empty.
	Peek() interface{}
}
