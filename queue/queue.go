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
)

// Interface is a type of Queue, which is FIFO(first-in-first-out).
type Interface interface {
	container.Interface

	// PushMulBack inserts an element into the tail of this Queue.
	Add(items ...interface{})
	// Peek retrieves, but does not remove, the head of this Queue, or return nil if this Queue is empty.
	Peek() interface{}
	// Poll retrieves and removes the head of the this Queue, or return nil if this Queue is empty.
	Poll() interface{}
}

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
}

var _ Interface = (*Queue)(nil)

// New creates a Queue.
func New() *Queue {
	return &Queue{}
}

func (sf *Queue) Len() int {
	return sf.length
}

// IsEmpty returns true if this Queue contains no elements.
func (sf *Queue) IsEmpty() bool {
	return sf.Len() == 0
}

// PushMulBack items to the queue
func (sf *Queue) Add(items ...interface{}) {
	for _, item := range items {
		e := element{
			next:  nil,
			value: item,
		}

		if sf.tail == nil {
			sf.head, sf.tail = &e, &e
		} else {
			sf.tail.next = &e
			sf.tail = &e
		}

		sf.length++
	}
}

func (sf *Queue) Peek() interface{} {
	if sf.head != nil {
		return sf.head.value
	}
	return nil
}

func (sf *Queue) Poll() interface{} {
	if sf.head != nil {
		val := sf.head.value

		sf.head = sf.head.next
		if nil == sf.head {
			sf.tail = nil
		}
		sf.length--

		return val
	}

	return nil
}

// Clear initializes or clears queue.
func (sf *Queue) Clear() {
	sf.head, sf.tail, sf.length = nil, nil, 0
}
