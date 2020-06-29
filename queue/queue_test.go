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

package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueueLen(t *testing.T) {
	q := New()
	q.Add(5)
	q.Add(6)
	assert.Equal(t, 2, q.Len())
}

func TestQueuePeek(t *testing.T) {
	q := New()
	q.Add(5)
	q.Add("hello")

	val1, ok := q.Peek().(int)
	assert.True(t, ok)
	assert.Equal(t, 5, val1)

	val2, ok := q.Peek().(int)
	assert.True(t, ok)
	assert.Equal(t, 5, val2)

	q.Poll()
	q.Poll()

	val3 := q.Peek()
	assert.Nil(t, val3)
}

func TestQueuePoll(t *testing.T) {
	q := New()

	q.Add(5)
	q.Add("hello")
	val1, ok := q.Poll().(int)
	assert.True(t, ok)
	assert.Equal(t, 5, val1)

	val2, ok := q.Poll().(string)
	assert.True(t, ok)
	assert.Equal(t, "hello", val2)

	val3 := q.Poll()
	assert.Nil(t, val3)
}

func TestQueueIsEmpty(t *testing.T) {
	q := New()
	q.Add(5)
	q.Add(6)
	assert.False(t, q.IsEmpty())

	q.Clear()
	assert.Equal(t, 0, q.Len())
	assert.True(t, q.IsEmpty())
}

func TestQueueInit(t *testing.T) {
	q := New()
	q.Add(5)
	q.Add(6)
	q.Clear()

	assert.Equal(t, 0, q.Len())
}
