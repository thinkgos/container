package comparator

import (
	"sort"
)

type sortContainer struct {
	items   []interface{}
	cmp     Comparator
	reverse bool
}

func (sf *sortContainer) Len() int {
	return len(sf.items)
}

func (sf *sortContainer) Swap(i, j int) {
	sf.items[i], sf.items[j] = sf.items[j], sf.items[i]
}

func (sf *sortContainer) Less(i, j int) bool {
	if sf.reverse {
		i, j = j, i
	}

	if nil != sf.cmp {
		return sf.cmp.Compare(sf.items[i], sf.items[j]) < 0
	}
	return Compare(sf.items[i], sf.items[j]) < 0
}

// Sort sorts values into ascending sequence according to their natural ordering, or according to the provided comparator.
func Sort(values []interface{}, c Comparator, reverse ...bool) {
	rev := false
	if len(reverse) > 0 {
		rev = reverse[0]
	}
	sort.Sort(&sortContainer{values, c, rev})
}
