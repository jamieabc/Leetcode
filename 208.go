package main

// Implement a trie with insert, search, and startsWith methods.
//
// Example:
//
// Trie trie = new Trie();
//
// trie.insert("apple");
// trie.search("apple");   // returns true
// trie.search("app");     // returns false
// trie.startsWith("app"); // returns true
// trie.insert("app");
// trie.search("app");     // returns true
// Note:
//
// You may assume that all inputs are consist of lowercase letters a-z.
// All inputs are guaranteed to be non-empty strings.

type Trie struct {
	IsWord   bool
	Children [26]*Trie
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	t := this

	for i := range word {
		if t.Children[word[i]-'a'] == nil {
			t.Children[word[i]-'a'] = &Trie{}
		}

		t = t.Children[word[i]-'a']
	}

	t.IsWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	t := this.Children[word[0]-'a']
	var i int

	for i = 1; i < len(word) && t != nil; i++ {
		t = t.Children[word[i]-'a']
	}

	return i == len(word) && t != nil && t.IsWord
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	t := this.Children[prefix[0]-'a']
	var i int

	for i = 1; i < len(prefix) && t != nil; i++ {
		t = t.Children[prefix[i]-'a']
	}

	return i == len(prefix) && t != nil
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
