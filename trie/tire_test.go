package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Simply make sure creating a new tree works.
func TestNewTrie(t *testing.T) {
	tries := NewTrie()

	require.Equal(t, 0, tries.Len())
	require.Equal(t, tries.root, tries.Root())
}

// Ensure that we can insert new keys into the tree, then check the size.
func TestInsert(t *testing.T) {
	tries := NewTrie()

	// We need to have an empty tree to begin with.
	require.Equal(t, 0, tries.Len())

	tries.Insert("key")
	tries.Insert("keys")

	// After inserting, we should have a size of two.
	require.Equal(t, 2, tries.Len())
}

// Ensure that MatchPrefix gives us the correct two keys in the tree.
func TestPrefixMatch(t *testing.T) {
	tries := NewTrie()

	// Feed it some fodder: only 'minio' and 'miny-os' should trip the matcher.
	tries.Insert("abcde")
	tries.Insert("hello")
	tries.Insert("kafka")
	tries.Insert("abcdef-os")

	matches := tries.MatchPrefix("abc")
	require.Equal(t, 2, len(matches))
	require.Contains(t, matches, "abcde")
	require.Contains(t, matches, "abcdef-os")

	matches = tries.MatchPrefix("invalid")
	require.Equal(t, 0, len(matches))

	matches = tries.MatchPrefix("kafkaa")
	require.Equal(t, 0, len(matches))

	matches = tries.MatchPrefix("hello")
	require.Equal(t, 1, len(matches))
}
