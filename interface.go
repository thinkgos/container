package container

// Stack is a Stack interface, which is LIFO (last-in-first-out).
type Stack interface {
	// Len returns the number of elements in the collection.
	Len() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear initializes or clears all of the elements from this container.
	Clear()
	// Push pushes an element into this Stack.
	Push(interface{})
	// Pop pops the element on the top of this Stack.
	Pop() interface{}
	// Peek retrieves, but does not remove, the element on the top of this Stack, or return nil if this Stack is empty.
	Peek() interface{}
}

// Queue is a type of Queue, which is FIFO(first-in-first-out).
type Queue interface {
	// Len returns the number of elements in the collection.
	Len() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear initializes or clears all of the elements from this container.
	Clear()
	// Add inserts an element into the tail of this Queue.
	Add(interface{})
	// Peek retrieves, but does not remove, the head of this Queue, or return nil if this Queue is empty.
	Peek() interface{}
	// Poll retrieves and removes the head of the this Queue, or return nil if this Queue is empty.
	Poll() interface{}
}

// PriorityQueue is a type of priority queue, and Queue implement this interface.
type PriorityQueue interface {
	Queue
	// Contains returns true if this queue contains the specified element.
	Contains(val interface{}) bool
	// Remove a single instance of the specified element from this queue, if it is present.
	// It returns false if the target value isn't present, otherwise returns true.
	Remove(val interface{})
}

// List is a type of list, both ArrayList and LinkedList implement this interface.
type List interface {
	// Len returns the number of elements in the collection.
	Len() int
	// IsEmpty returns true if this container contains no elements.
	IsEmpty() bool
	// Clear initializes or clears all of the elements from this container.
	Clear()

	// Push appends the specified elements to the end of this list.
	Push(vals interface{})
	// PushFront inserts a new element e with value v at the front of list l
	PushFront(v interface{})
	// PushBack inserts a new element e with value v at the back of list l.
	PushBack(v interface{})
	// Add inserts the specified element at the specified position in this list.
	Add(index int, val interface{}) error

	// Poll return the front element value and then remove from list
	Poll() interface{}
	// PollFront return the front element value and then remove from list
	PollFront() interface{}
	// PollBack return the back element value and then remove from list
	PollBack() interface{}
	// Remove removes the element at the specified position in this list.
	// It returns an error if the index is out of range.
	Remove(index int) (interface{}, error)
	// RemoveWithValue removes the first occurence of the specified element from this list, if it is present.
	// It returns false if the target value isn't present, otherwise returns true.
	RemoveWithValue(val interface{}) bool

	// Get returns the element at the specified positon in this list. The index must be in the range of [0, size).
	Get(index int) (interface{}, error)
	// Peek return the front element value
	Peek() interface{}
	// PeekFront return the front element value
	PeekFront() interface{}
	// PeekBack return the back element value
	PeekBack() interface{}

	// Iterator returns an iterator over the elements in this list in proper sequence.
	Iterator(f func(interface{}) bool)
	// ReverseIterator returns an iterator over the elements in this list in reverse sequence as Iterator.
	ReverseIterator(f func(interface{}) bool)

	// Contains returns true if this list contains the specified element.
	Contains(val interface{}) bool
	// Sort sorts the element using default options below. It sorts the elements into ascending sequence according to their natural ordering.
	Sort(reverse ...bool)
	// Values get a copy of all the values in the list
	Values() []interface{}
}
