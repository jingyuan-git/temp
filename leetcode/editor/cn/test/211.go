package main

type TrieNode struct {
	children map[byte]*TrieNode
	isEnd    bool
}

func newTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
		isEnd:    false,
	}
}

type WordDictionary struct {
	root *TrieNode
}

/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: newTrieNode(),
	}
}

/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		if _, ok := node.children[word[i]]; !ok {
			node.children[word[i]] = newTrieNode()
		}
		node = node.children[word[i]]
	}
	node.isEnd = true
}

/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {
	node := this.root
	return search(word, node)
}

func search(word string, node *TrieNode) bool {
	if node == nil {
		return false
	}
	if word == "" {
		return node.isEnd
	}

	if word[0] == '.' {
		for _, n := range node.children {
			if search(word[1:], n) {
				return true
			}
		}
	}
	node = node.children[word[0]]
	return search(word[1:], node)
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
//
//作者：bx0216
//链接：https://leetcode-cn.com/problems/design-add-and-search-words-data-structure/solution/trieshi-xian-by-bx0216/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
