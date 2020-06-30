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

package queue

import (
	"github.com/thinkgos/container"
	"github.com/thinkgos/container/comparator"
)

// element is an element of the Queue
type element struct {
	next  *element
	value interface{}
}

// Queue represents a singly linked list.
type Queue struct {
	head   *element
	tail   *element
	length int
	cmp    comparator.Comparator
}

// Option option for New
type Option func(q *Queue)

var _ container.Queue = (*Queue)(nil)

// WithComparator with user's Comparator
func WithComparator(c comparator.Comparator) Option {
	return func(q *Queue) {
		q.cmp = c
	}
}

// New creates a Queue. which implement queue.Interface
func New(opts ...Option) *Queue {
	q := new(Queue)
	for _, opt := range opts {
		opt(q)
	}
	return q
}

// Len returns the length of this priority queue.
func (sf *Queue) Len() int { return sf.length }

// IsEmpty returns true if this Queue contains no elements.
func (sf *Queue) IsEmpty() bool { return sf.Len() == 0 }

// Clear initializes or clears queue.
func (sf *Queue) Clear() { sf.head, sf.tail, sf.length = nil, nil, 0 }

// Add items to the queue
func (sf *Queue) Add(v interface{}) {
	e := &element{value: v}
	if sf.tail == nil {
		sf.head, sf.tail = e, e
	} else {
		sf.tail.next = e
		sf.tail = e
	}
	sf.length++
}

// Peek retrieves, but does not remove, the head of this Queue, or return nil if this Queue is empty.
func (sf *Queue) Peek() interface{} {
	if sf.head != nil {
		return sf.head.value
	}
	return nil
}

// Poll retrieves and removes the head of the this Queue, or return nil if this Queue is empty.
func (sf *Queue) Poll() interface{} {
	var val interface{}

	if sf.head != nil {
		val = sf.head.value
		sf.head = sf.head.next
		if sf.head == nil {
			sf.tail = nil
		}
		sf.length--
	}
	return val
}

// Contains returns true if this queue contains the specified element.
func (sf *Queue) Contains(val interface{}) bool {
	for e := sf.head; e != nil; e = e.next {
		if sf.compare(val, e.value) {
			return true
		}
	}
	return false
}

// Remove a single instance of the specified element from this queue, if it is present.
func (sf *Queue) Remove(val interface{}) {
	for pre, e := sf.head, sf.head; e != nil; {
		if sf.compare(val, e.value) {
			if sf.head == e && sf.tail == e {
				sf.head, sf.tail = nil, nil
			} else if sf.head == e {
				sf.head = e.next
			} else if sf.tail == e {
				sf.tail = pre
				sf.tail.next = nil
			} else {
				pre.next = e.next
			}
			e.next = nil
			sf.length--
			return
		}
		pre = e
		e = e.next
	}
}

func (sf *Queue) compare(v1, v2 interface{}) bool {
	if sf.cmp != nil {
		return sf.cmp.Compare(v1, v2) == 0
	}
	return v1 == v2
}
