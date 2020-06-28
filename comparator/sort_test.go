package comparator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 9, 15, 19}
	assertSort(t, input1, expected1, false, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "john", "roy", "tom"}
	assertSort(t, input2, expected2, false, nil)
}

func TestSortWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{19, 15, 9, 6, 4}
	assertSort(t, input1, expected1, false, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"tom", "roy", "john", "benjamin", "alice"}
	assertSort(t, input2, expected2, false, reverseString{})
}

func TestReverseSort(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{19, 15, 9, 6, 4}
	assertSort(t, input1, expected1, true, nil)

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"tom", "roy", "john", "benjamin", "alice"}
	assertSort(t, input2, expected2, true, nil)
}

func TestReverseSortWithComparator(t *testing.T) {
	input1 := []interface{}{6, 4, 9, 19, 15}
	expected1 := []interface{}{4, 6, 9, 15, 19}
	assertSort(t, input1, expected1, true, reverseInt{})

	input2 := []interface{}{"benjamin", "alice", "john", "tom", "roy"}
	expected2 := []interface{}{"alice", "benjamin", "john", "roy", "tom"}
	assertSort(t, input2, expected2, true, reverseString{})
}

func assertSort(t *testing.T, input []interface{}, expected []interface{}, reverse bool, c Comparator) {
	// sort
	Sort(input, c, reverse)
	for i := 0; i < len(input); i++ {
		assert.Equal(t, expected[i], input[i])
	}
}

type reverseString struct{}

// Compare returns reverse order for string
func (i reverseString) Compare(v1, v2 interface{}) int {
	i1, i2 := v1.(string), v2.(string)

	if i1 < i2 {
		return 1
	}
	if i1 > i2 {
		return -1
	}
	return 0
}

type reverseInt struct{}

// Compare returns reverse order for int
func (i reverseInt) Compare(v1, v2 interface{}) int {
	i1, i2 := v1.(int), v2.(int)

	if i1 < i2 {
		return 1
	}
	if i1 > i2 {
		return -1
	}
	return 0
}
