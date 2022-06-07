package DataStructures

// AlphabetSize for the language used
const AlphabetSize = 26

// TrieNode Structure for use with Children notes
type TrieNode struct {
	children [AlphabetSize]*TrieNode
	isEnd    bool
}

// Trie Structure
type Trie struct {
	root *TrieNode
}

// Insert word runes into tree
func (t *Trie) Insert(s string) {
	c := t.root
	for _, r := range s {
		i := r - 'a'
		if c.children[i] == nil {
			c.children[i] = &TrieNode{}
		}
		c = c.children[i]
	}

	c.isEnd = true
}

// Search takes a string and checkes to see if it exists
// returns true if found otherwise false
func (t *Trie) Search(s string) bool {
	c := t.root
	for _, r := range s {
		i := r - 'a'
		if c.children[i] == nil {
			return false
		}
		c = c.children[i]
	}

	if c.isEnd {
		return true
	}

	return false
}

// NewTrie creates a new Trie to be used.
func NewTrie() *Trie {
	return &Trie{root: &TrieNode{}}
}
