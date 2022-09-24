package main

type Trie struct {
	children	map[byte]*Trie
	isEnd		bool
}

func NewTrie() Trie {
	return Trie{children: map[byte]*Trie{}}
}

func (t *Trie) Insert(word string) {
	node := t
	for i := range word {
		if _, ok := node.children[word[i]]; !ok {
			node.children[word[i]] = &Trie{children: map[byte]*Trie{}}
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

func (t *Trie) SearchPrefix(prefix string) *Trie {
	node := t
	for i := range prefix {
		if _, ok := node.children[prefix[i]]; !ok {
			return nil
		}
		node = node.children[prefix[i]]
	}
	return node
}

func (t *Trie) Search(word string) bool {
	node := t.SearchPrefix(word)
	return node != nil && node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.SearchPrefix(prefix) != nil
}
