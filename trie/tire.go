// Copyright [2020] [thinkgos] thinkgo@aliyun.com
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

package trie

// Node trie tree node container carries value and children.
type Node struct {
	exist bool
	value interface{}
	child map[rune]*Node // runes as child.
}

// newNode create a new trie node.
func newNode() *Node {
	return &Node{
		false,
		nil,
		make(map[rune]*Node),
	}
}

// Trie is a trie container.
type Trie struct {
	root *Node
	size int
}

// NewTrie create a new trie.
func NewTrie() *Trie {
	return &Trie{newNode(), 0}
}

// Root returns root node.
func (t *Trie) Root() *Node { return t.root }

// Len returns trie size.
func (t *Trie) Len() int { return t.size }

// Insert insert a key.
func (t *Trie) Insert(key string) {
	cur := t.root
	for _, v := range key {
		if cur.child[v] == nil {
			cur.child[v] = newNode()
		}
		cur = cur.child[v]
	}

	if !cur.exist {
		// increment when new rune child is added.
		t.size++
		cur.exist = true
	}
	// value is stored for retrieval in future.
	cur.value = key
}

// MatchPrefix - prefix match.
func (t *Trie) MatchPrefix(key string) []interface{} {
	node, _ := t.findNode(key)
	if node != nil {
		return t.Walk(node)
	}
	return []interface{}{}
}

// Walk the tree, return the trie node value slice
func (t *Trie) Walk(node *Node) (ret []interface{}) {
	if node.exist {
		ret = append(ret, node.value)
	}
	for _, v := range node.child {
		ret = append(ret, t.Walk(v)...)
	}
	return ret
}

// find nodes corresponding to key.
func (t *Trie) findNode(key string) (node *Node, index int) {
	cur := t.root
	b := false
	for k, v := range key {
		if b {
			index = k
			b = false
		}
		if cur.child[v] == nil {
			return nil, index
		}
		cur = cur.child[v]
		if cur.exist {
			b = true
		}
	}

	if cur.exist {
		index = len(key)
	}
	return cur, index
}
