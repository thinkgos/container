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

// Package linked implements both an arrayList and a List.
package linked

import (
	"container/list"
	"fmt"

	"github.com/thinkgos/container/comparator"
)

// List represents a doubly linked list.
// It implements the interface list.Interface.
type List struct {
	l   *list.List
	cmp comparator.Comparator
}

type Option func(l *List)

func WithComparator(cmp comparator.Comparator) Option {
	return func(l *List) {
		l.cmp = cmp
	}
}

// New initializes and returns an List.
func New(opts ...Option) *List {
	l := &List{l: list.New()}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (sf *List) Clear() {
	sf.l.Init()
}

func (sf *List) Len() int {
	return sf.l.Len()
}

func (sf *List) PushFront(items ...interface{}) {
	for _, item := range items {
		sf.l.PushFront(item)
	}
}

func (sf *List) PushBack(items ...interface{}) {
	for _, item := range items {
		sf.l.PushBack(item)
	}
}

func (sf *List) PushFrontList(other *List) {
	sf.l.PushFrontList(other.l)
}

func (sf *List) PushBackList(other *List) {
	sf.l.PushBackList(other.l)
}

func (sf *List) InsertBefore(v interface{}, mark *list.Element) {
	sf.l.InsertBefore(v, mark)
}

func (sf *List) InsertAfter(v interface{}, mark *list.Element) {
	sf.l.InsertAfter(v, mark)
}

func (sf *List) MoveToFront(e *list.Element) {
	sf.l.MoveToFront(e)
}

func (sf *List) MoveToBack(e *list.Element) {
	sf.l.MoveToBack(e)
}

func (sf *List) MoveBefore(e, mark *list.Element) {
	sf.l.MoveBefore(e, mark)
}

func (sf *List) MoveAfter(e, mark *list.Element) {
	sf.l.MoveAfter(e, mark)
}
func (sf *List) Remove(e *list.Element) interface{} {
	return sf.l.Remove(e)
}

func (sf *List) IsEmpty() bool {
	return sf.l.Len() == 0
}

func (sf *List) AddTo(index int, val interface{}) error {
	length := sf.Len()
	if index < 0 || index > length {
		return fmt.Errorf("Index out of range, index: %d, len: %d", index, length)
	}

	if index == length {
		sf.PushBack(val)
	} else {
		sf.InsertBefore(val, sf.getElement(index))
	}
	return nil
}

// getElement returns the element at the specified positon.
func (sf *List) getElement(index int) *list.Element {
	var i int
	var e *list.Element

	if length := sf.Len(); index < (length >> 1) {
		for i, e = 0, sf.l.Front(); i < index; i, e = i+1, e.Next() {
		}
	} else {
		for i, e = length-1, sf.l.Back(); i > index; i, e = i-1, e.Prev() {
		}
	}
	return e
}

// indexOf returns the index of the first occurence of the specified element
// in this list, or -1 if this list does not contain the element.
func (sf *List) indexOf(val interface{}) int {
	index := 0

	for e := sf.l.Front(); e != nil; e = e.Next() {
		if sf.compare(val, e.Value) {
			return index
		}
		index++
	}

	return -1
}

func (sf *List) Contains(val interface{}) bool {
	return sf.indexOf(val) >= 0
}

func (sf *List) Get(index int) (interface{}, error) {
	length := sf.Len()
	if index < 0 || index >= length {
		return nil, fmt.Errorf("Index out of range, index:%d, len:%d", index, length)
	}

	return sf.getElement(index).Value, nil
}

func (sf *List) RemoveWithIndex(index int) (interface{}, error) {
	size := sf.Len()
	if index < 0 || index >= size {
		return nil, fmt.Errorf("Index out of range, index:%d, len:%d", index, size)
	}
	return sf.Remove(sf.getElement(index)), nil
}

func (sf *List) RemoveWithValue(val interface{}) bool {
	if sf.Len() == 0 {
		return false
	}
	for e := sf.l.Front(); e != nil; e = e.Next() {
		if sf.compare(val, e.Value) {
			sf.Remove(e)
			return true
		}
	}
	return false
}

func (sf *List) Iterator(cb func(interface{}) bool) {
	for e := sf.l.Front(); e != nil; e = e.Next() {
		if cb == nil || !cb(e.Value) {
			return
		}
	}
}

func (sf *List) ReverseIterator(cb func(interface{}) bool) {
	for e := sf.l.Back(); e != nil; e = e.Prev() {
		if cb == nil || !cb(e.Value) {
			return
		}
	}
}

func (sf *List) compare(v1, v2 interface{}) bool {
	if sf.cmp != nil {
		return sf.cmp.Compare(v1, v2) == 0
	}
	return v1 == v2
}

func (sf *List) Sort(reverse ...bool) {
	if sf.Len() < 2 {
		return
	}

	// get all the values
	vals := sf.values()

	// sort the data
	comparator.Sort(vals, sf.cmp, reverse...)

	// clear the linked list
	sf.Clear()

	// add the sorted values into the list again
	sf.PushBack(vals...)
}

func (sf *List) values() []interface{} {
	if sf.Len() == 0 {
		return []interface{}{}
	}

	values := make([]interface{}, 0, sf.Len())

	sf.Iterator(func(v interface{}) bool {
		values = append(values, v)
		return true
	})

	return values
}
