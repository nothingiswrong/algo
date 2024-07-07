package trie



type trieNode struct {
	children map[rune] *trieNode
	terminate bool
}

type Trie struct {
	root *trieNode
}


func New() *Trie {
	trie := new(Trie)
	trie.root = &trieNode{make(map[rune]*trieNode), false}
	return trie
}

func (t *Trie) Add(s string) {
	it := t.root
	for _, r := range s {
		_, ok := it.children[r]
		if !ok {
			it.children[r] = &trieNode{make(map[rune]*trieNode), false}
		}
		it = it.children[r]
	}
	it.terminate = true
}

func (t *Trie) Consists(s string) bool {
	it := t.root
	for _, r := range s {
		if _, ok := it.children[r]; !ok {
			return false
		}
		it = it.children[r]
	}
	if !it.terminate {
		return false
	}
	return true
}

func (t *Trie) Remove(s string) bool {
	return false

}
