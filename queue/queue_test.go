// Copyright (c) 2019, Benjamin Wang (benjamin_wang@aliyun.com). All rights reserved.
// Licensed under the MIT license that can be found in the LICENSE file.

package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueSize(t *testing.T) {
	q := New()
	q.Add(5, 6)

	assert.Equal(t, 2, q.Len())
}

func TestQueuePeek(t *testing.T) {
	q := New()
	q.Add(5, "hello")

	val1, ok := q.Peek().(int)
	assert.True(t, ok)
	assert.Equal(t, 5, val1)

	val2, ok := q.Peek().(int)
	assert.True(t, ok)
	assert.Equal(t, 5, val2)
}

func TestQueuePoll(t *testing.T) {
	q := New()

	q.Add(5, "hello")

	val1, ok := q.Poll().(int)
	assert.True(t, ok)
	assert.Equal(t, 5, val1)

	val2, ok := q.Poll().(string)
	assert.True(t, ok)
	assert.Equal(t, "hello", val2)
}

func TestQueueIsEmpty(t *testing.T) {
	q := New()
	q.Add(5, 6)
	assert.False(t, q.IsEmpty())

	q.Init()
	assert.True(t, q.IsEmpty())
}

func TestQueueInit(t *testing.T) {
	q := New()

	q.Add(5, 6)
	q.Init()

	assert.Equal(t, 0, q.Len())
}
