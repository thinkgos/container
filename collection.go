package container

// Interface is a type of collection, all containers should implement this interface.
type Interface interface {
	// Len returns the number of elements in the collection.
	Len() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear initializes or clears all of the elements from this container.
	Clear()
}
