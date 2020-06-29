package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickStackLen(t *testing.T) {
	s := NewQuickStack()
	s.Push(5)
	s.Push(6)
	assert.Equal(t, 2, s.Len())
}

func TestQuickStackValue(t *testing.T) {
	s := NewQuickStack()
	s.Push(5)
	s.Push("hello")

	// Peek "hello"
	val1, ok := s.Peek().(string)
	assert.True(t, ok)
	assert.Equal(t, "hello", val1)

	// Pop "hello"
	val2, ok := s.Pop().(string)
	assert.True(t, ok)
	assert.Equal(t, "hello", val2)

	// Peek 5
	val3, ok := s.Peek().(int)
	assert.True(t, ok)
	assert.Equal(t, 5, val3)

	// Pop 5
	val4, ok := s.Pop().(int)
	assert.True(t, ok)
	assert.Equal(t, 5, val4)

	val5 := s.Pop()
	assert.Nil(t, val5)
	assert.Nil(t, s.Peek())
}

func TestQuickStackIsEmpty(t *testing.T) {
	s := NewQuickStack()
	s.Push(5)
	s.Push(6)

	assert.False(t, s.IsEmpty())
	s.Clear()
	assert.True(t, s.IsEmpty())
}

func TestQuickStackInit(t *testing.T) {
	s := NewQuickStack()
	s.Push(5)
	s.Push(6)

	s.Clear()
	assert.Zero(t, s.Len())
}
