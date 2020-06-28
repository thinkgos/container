package comparator

// Comparator imposes a total ordering on some collection of objects, and it allows precise control over the sort order.
type Comparator interface {
	// Compare compares its two arguments for order.
	// It returns a negative integer, zero, or a positive integer as the first argument is less than, equal to, or greater than the second.
	Compare(v1, v2 interface{}) int
}
