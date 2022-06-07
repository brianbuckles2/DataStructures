package DataStructures

import (
	"testing"
)

func TestTrieInsert(t *testing.T) {
	trie := NewTrie()
	w := "iron"
	trie.Insert(w)

	if trie.root.children[8] == nil {
		t.Fatalf(`Trie insert = %v want %s`, trie.root.children[8], "i")
	}

	if trie.root.children[8].children[17] == nil {
		t.Fatalf(`Trie insert = %v want %s`, trie.root.children[8].children[17], "r")
	}

	if trie.root.children[8].children[17].children[14] == nil {
		t.Fatalf(`Trie insert = %v want %s`, trie.root.children[8].children[17].children[14], "o")
	}

	if trie.root.children[8].children[17].children[14].children[13] == nil {
		t.Fatalf(`Trie insert = %v want %s`, trie.root.children[8].children[17].children[14].children[13], "n")
	}

	if !trie.root.children[8].children[17].children[14].children[13].isEnd {
		t.Fatalf(`Trie insert = %v want %s`, trie.root.children[8].children[17].children[14].children[13].isEnd, "true")
	}
}

func TestTrieSearch(t *testing.T) {
	trie := NewTrie()
	words := [5]string{"iron", "ironman", "test", "ir", "cupcake"}

	for _, w := range words {
		trie.Insert(w)
	}

	if !trie.Search("iron") {
		t.Fatalf("iron was not found")
	}
}
