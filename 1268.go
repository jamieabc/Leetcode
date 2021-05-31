package main

// Given an array of strings products and a string searchWord. We want to design a system that suggests at most three product names from products after each character of searchWord is typed. Suggested products should have common prefix with the searchWord. If there are more than three products with a common prefix return the three lexicographically minimums products.
//
// Return list of lists of the suggested products after each character of searchWord is typed.
//
//
//
// Example 1:
//
// Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
// Output: [
// ["mobile","moneypot","monitor"],
// ["mobile","moneypot","monitor"],
// ["mouse","mousepad"],
// ["mouse","mousepad"],
// ["mouse","mousepad"]
// ]
// Explanation: products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"]
// After typing m and mo all products match and we show user ["mobile","moneypot","monitor"]
// After typing mou, mous and mouse the system suggests ["mouse","mousepad"]
//
// Example 2:
//
// Input: products = ["havana"], searchWord = "havana"
// Output: [["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
//
// Example 3:
//
// Input: products = ["bags","baggage","banner","box","cloths"], searchWord = "bags"
// Output: [["baggage","bags","banner"],["baggage","bags","banner"],["baggage","bags"],["bags"]]
//
// Example 4:
//
// Input: products = ["havana"], searchWord = "tatiana"
// Output: [[],[],[],[],[],[],[]]
//
//
//
// Constraints:
//
// 1 <= products.length <= 1000
// There are no repeated elements in products.
// 1 <= Σ products[i].length <= 2 * 10^4
// All characters of products[i] are lower-case English letters.
// 1 <= searchWord.length <= 1000
// All characters of searchWord are lower-case English letters.

type Trie struct {
	IsWord   bool
	Children [26]*Trie
}

func (t *Trie) Insert(word string) {
	node := t

	for i := 0; i < len(word); i++ {
		if node.Children[word[i]-'a'] == nil {
			node.Children[word[i]-'a'] = &Trie{}
		}

		node = node.Children[word[i]-'a']
	}
	node.IsWord = true
}

func (t *Trie) Dfs(prefix []byte, ans *[]string, count *int) {
	if t == nil {
		return
	}

	if t.IsWord {
		*ans = append(*ans, string(prefix))
		*count--
	}

	for i := 0; i < len(t.Children) && *count > 0; i++ {
		if t.Children[i] != nil {
			tmp := append([]byte{}, prefix...)
			tmp = append(tmp, byte('a'+i))

			t.Children[i].Dfs(tmp, ans, count)
		}
	}
}

func suggestedProducts(products []string, searchWord string) [][]string {
	root := &Trie{}

	for _, product := range products {
		root.Insert(product)
	}

	ans := make([][]string, 0)
	size := len(searchWord)
	var count int

	for i, node := 0, root; i < size; i++ {
		candidates := make([]string, 0)

		if node != nil {
			node = node.Children[searchWord[i]-'a']
			count = 3

			node.Dfs([]byte(searchWord[:i+1]), &candidates, &count)
		}

		ans = append(ans, candidates)
	}

	return ans
}
