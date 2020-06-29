package stack

// QuickStack is quick LIFO stack
type QuickStack struct {
	items []interface{}
}

var _ Interface = (*QuickStack)(nil)

// New creates a QuickStack. which implement interface stack.Interface
func NewQuickStack() *QuickStack { return &QuickStack{} }

// Len returns the length of this priority queue.
func (s *QuickStack) Len() int { return len(s.items) }

// IsEmpty returns true if this QuickStack contains no elements.
func (s *QuickStack) IsEmpty() bool { return len(s.items) == 0 }

// Clear removes all the elements from this QuickStack.
func (s *QuickStack) Clear() { s.items = nil }

// Push pushes an element into this QuickStack.
func (s *QuickStack) Push(val interface{}) { s.items = append(s.items, val) }

// Pop pops the element on the top of this QuickStack.
func (s *QuickStack) Pop() interface{} {
	if length := len(s.items); length > 0 {
		val := s.items[length-1]
		s.items = s.items[:length-1]
		return val
	}
	return nil
}

// Peek retrieves, but does not remove, the element on the top of this QuickStack, or return nil if this QuickStack is empty.
func (s *QuickStack) Peek() interface{} {
	if len(s.items) > 0 {
		return s.items[len(s.items)-1]
	}
	return nil
}
