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

package comparator

import (
	"container/heap"
	"sort"
)

var (
	_ heap.Interface = (*Container)(nil)
	_ sort.Interface = (*Container)(nil)
)

// Container for sort or heap sort, it implement sort.Interface and heap.Interface.
type Container struct {
	noCopy  noCopy // nolint: structcheck,unused
	Items   []interface{}
	Cmp     Comparator
	Reverse bool
}

// Len implement heap.Interface.
func (sf *Container) Len() int {
	return len(sf.Items)
}

// Swap implement heap.Interface.
func (sf *Container) Swap(i, j int) {
	sf.Items[i], sf.Items[j] = sf.Items[j], sf.Items[i]
}

// Less implement heap.Interface.
func (sf *Container) Less(i, j int) bool {
	if sf.Reverse {
		i, j = j, i
	}

	if sf.Cmp != nil {
		return sf.Cmp.Compare(sf.Items[i], sf.Items[j]) < 0
	}
	return Compare(sf.Items[i], sf.Items[j]) < 0
}

// Push implement heap.Interface.
func (sf *Container) Push(x interface{}) {
	sf.Items = append(sf.Items, x)
}

// Pop implement heap.Interface.
func (sf *Container) Pop() interface{} {
	old := sf.Items
	n := len(old)
	x := old[n-1]
	old[n-1] = nil // should set nil for gc
	sf.Items = old[:n-1]
	return x
}

// Sort sorts values into ascending sequence according to their natural ordering,
// or according to the provided comparator.
func (sf *Container) Sort() {
	sort.Sort(sf)
}

// Sort sorts values into ascending sequence according to their natural ordering,
// or according to the provided comparator.
func Sort(values []interface{}, c Comparator, reverse ...bool) {
	rev := false
	if len(reverse) > 0 {
		rev = reverse[0]
	}
	sort.Sort(&Container{Items: values, Cmp: c, Reverse: rev})
}
