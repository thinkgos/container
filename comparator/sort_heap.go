package comparator

import (
	"container/heap"
	"sort"
)

var (
	_ heap.Interface = (*Container)(nil)
	_ sort.Interface = (*Container)(nil)
)

type Container struct {
	noCopy  noCopy
	Items   []interface{}
	Cmp     Comparator
	Reverse bool
}

func (sf *Container) Len() int {
	return len(sf.Items)
}

func (sf *Container) Swap(i, j int) {
	sf.Items[i], sf.Items[j] = sf.Items[j], sf.Items[i]
}

func (sf *Container) Less(i, j int) bool {
	if sf.Reverse {
		i, j = j, i
	}

	if nil != sf.Cmp {
		return sf.Cmp.Compare(sf.Items[i], sf.Items[j]) < 0
	}
	return Compare(sf.Items[i], sf.Items[j]) < 0
}

func (sf *Container) Push(x interface{}) {
	sf.Items = append(sf.Items, x)
}

func (sf *Container) Pop() interface{} {
	old := sf.Items
	n := len(old)
	x := old[n-1]
	sf.Items = old[:n-1]
	return x
}

func (sf *Container) Sort() {
	sort.Sort(sf)
}

// Sort sorts values into ascending sequence according to their natural ordering, or according to the provided comparator.
func Sort(values []interface{}, c Comparator, reverse ...bool) {
	rev := false
	if len(reverse) > 0 {
		rev = reverse[0]
	}
	sort.Sort(&Container{Items: values, Cmp: c, Reverse: rev})
}
