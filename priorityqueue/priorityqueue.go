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

// Package priorityqueue implements an unbounded priority queue based on a priority heap.
// The elements of the priority queue are ordered according to their natural ordering,
// or by a Comparator provided at PriorityQueue construction time.
package priorityqueue

import (
	"container/heap"

	"github.com/thinkgos/container"
	"github.com/thinkgos/container/comparator"
)

var _ container.Queue = (*Queue)(nil)

// Queue represents an unbounded priority queue based on a priority heap.
// It implements heap.Interface.
type Queue struct {
	ctn *comparator.Container
}

// Option option for New.
type Option func(q *Queue)

// WithComparator with user's Comparator.
func WithComparator(c comparator.Comparator) Option {
	return func(q *Queue) { q.ctn.Cmp = c }
}

// WithMaxHeap with max heap.
func WithMaxHeap(b bool) Option {
	return func(q *Queue) {
		q.ctn.Reverse = b
	}
}

// New initializes and returns an Queue, default min heap.
func New(opts ...Option) *Queue {
	q := &Queue{
		ctn: &comparator.Container{
			Items:   []interface{}{},
			Reverse: false,
		},
	}
	for _, opt := range opts {
		opt(q)
	}
	return q
}

// Len returns the length of this priority queue.
func (sf *Queue) Len() int { return sf.ctn.Len() }

// IsEmpty returns true if this list contains no elements.
func (sf *Queue) IsEmpty() bool { return sf.Len() == 0 }

// Clear removes all of the elements from this priority queue.
func (sf *Queue) Clear() { sf.ctn.Items = make([]interface{}, 0) }

// Add inserts the specified element into this priority queue.
func (sf *Queue) Add(items interface{}) {
	heap.Push(sf.ctn, items)
}

// Peek retrieves, but does not remove, the head of this queue, or return nil if this queue is empty.
func (sf *Queue) Peek() interface{} {
	if sf.Len() > 0 {
		return sf.ctn.Items[0]
	}
	return nil
}

// Poll retrieves and removes the head of the this queue, or return nil if this queue is empty.
func (sf *Queue) Poll() interface{} {
	if sf.Len() > 0 {
		return heap.Pop(sf.ctn)
	}
	return nil
}

// Contains returns true if this queue contains the specified element.
func (sf *Queue) Contains(val interface{}) bool { return sf.indexOf(val) >= 0 }

// Remove a single instance of the specified element from this queue, if it is present.
// It returns false if the target value isn't present, otherwise returns true.
func (sf *Queue) Remove(val interface{}) {
	if idx := sf.indexOf(val); idx >= 0 {
		heap.Remove(sf.ctn, idx)
	}
}

func (sf *Queue) indexOf(val interface{}) int {
	if sf.Len() > 0 && val != nil {
		for i := 0; i < sf.Len(); i++ {
			if sf.compare(val, sf.ctn.Items[i]) {
				return i
			}
		}
	}
	return -1
}

func (sf *Queue) compare(v1, v2 interface{}) bool {
	if sf.ctn.Cmp != nil {
		return sf.ctn.Cmp.Compare(v1, v2) == 0
	}
	return v1 == v2
}
