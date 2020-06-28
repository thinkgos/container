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

// Package stack implements a stack, which orders elements in a LIFO (last-in-first-out) manner.
package stack

import (
	"container/list"

	"github.com/thinkgos/container"
)

// Interface is a stack, which is LIFO (last-in-first-out).
type Interface interface {
	container.Interface

	// Push pushes an element into this stack.
	Push(val interface{})
	// Pop pops the element on the top of this stack.
	Pop() interface{}
	// Peek retrieves, but does not remove, the element on the top of this stack, or return nil if this stack is empty.
	Peek() interface{}
}

// stack is LIFO.
type stack struct {
	l *list.List
}

// New creates a stack.
func New() Interface {
	return &stack{list.New()}
}

func (s *stack) Len() int {
	return s.l.Len()
}

// IsEmpty returns true if this stack contains no elements.
func (s *stack) IsEmpty() bool {
	return s.l.Len() == 0
}

func (s *stack) Push(val interface{}) {
	s.l.PushFront(val)
}

func (s *stack) Pop() interface{} {
	if ele := s.l.Front(); ele != nil {
		return s.l.Remove(ele)

	}
	return nil
}

func (s *stack) Peek() interface{} {
	return s.l.Front().Value
}

// Clear removes all the elements from this stack.
func (s *stack) Init() {
	s.l.Init()
}
