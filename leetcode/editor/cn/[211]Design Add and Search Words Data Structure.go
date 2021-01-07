//Design a data structure that supports adding new words and finding if a string
// matches any previously added string. 
//
// Implement the WordDictionary class: 
//
// 
// WordDictionary() Initializes the object. 
// void addWord(word) Adds word to the data structure, it can be matched later. 
//
// bool search(word) Returns true if there is any string in the data structure t
//hat matches word or false otherwise. word may contain dots '.' where dots can be
// matched with any letter. 
// 
//
// 
// Example: 
//
// 
//Input
//["WordDictionary","addWord","addWord","addWord","search","search","search","se
//arch"]
//[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]
//Output
//[null,null,null,null,false,true,true,true]
//
//Explanation
//WordDictionary wordDictionary = new WordDictionary();
//wordDictionary.addWord("bad");
//wordDictionary.addWord("dad");
//wordDictionary.addWord("mad");
//wordDictionary.search("pad"); // return False
//wordDictionary.search("bad"); // return True
//wordDictionary.search(".ad"); // return True
//wordDictionary.search("b.."); // return True
// 
//
// 
// Constraints: 
//
// 
// 1 <= word.length <= 500 
// word in addWord consists lower-case English letters. 
// word in search consist of '.' or lower-case English letters. 
// At most 50000 calls will be made to addWord and search. 
// 
// Related Topics 设计 字典树 回溯算法 
// 👍 196 👎 0

package main
//leetcode submit region begin(Prohibit modification and deletion)
type trieNode struct {
	next 	map[byte]*trieNode
	isEnd	bool
}

func NewTrieNode() *trieNode {
	return &trieNode{
		next: make(map[byte]*trieNode),
		isEnd: false,
	}
}

type WordDictionary struct {
	root	*trieNode
}


/** Initialize your data structure here. */
func Constructor() WordDictionary {
	return WordDictionary{
		root: NewTrieNode(),
	}
}


/** Adds a word into the data structure. */
func (this *WordDictionary) AddWord(word string)  {
	node := this
	for _, v := range word {
		v -= 'a'
		if
	}
}


/** Returns if the word is in the data structure. A word could contain the dot character '.' to represent any one letter. */
func (this *WordDictionary) Search(word string) bool {

}


/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
//leetcode submit region end(Prohibit modification and deletion)
