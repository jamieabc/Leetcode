package main

//Given two sentences words1, words2 (each represented as an array of strings), and a list of similar word pairs pairs, determine if two sentences are similar.
//
//For example, "great acting skills" and "fine drama talent" are similar, if the similar word pairs are pairs = [["great", "fine"], ["acting","drama"], ["skills","talent"]].
//
//Note that the similarity relation is not transitive. For example, if "great" and "fine" are similar, and "fine" and "good" are similar, "great" and "good" are not necessarily similar.
//
//However, similarity is symmetric. For example, "great" and "fine" being similar is the same as "fine" and "great" being similar.
//
//Also, a word is always similar with itself. For example, the sentences words1 = ["great"], words2 = ["great"], pairs = [] are similar, even though there are no specified similar word pairs.
//
//Finally, sentences can only be similar if they have the same number of words. So a sentence like words1 = ["great"] can never be similar to words2 = ["doubleplus","good"].
//
//Note:
//
//    The length of words1 and words2 will not exceed 1000.
//    The length of pairs will not exceed 2000.
//    The length of each pairs[i] will be 2.
//    The length of each words[i] and pairs[i][j] will be in the range [1, 20].

func areSentencesSimilar(words1 []string, words2 []string, pairs [][]string) bool {
	mapping := make(map[string]map[string]bool)
	for _, p := range pairs {
		if _, ok := mapping[p[0]]; !ok {
			mapping[p[0]] = make(map[string]bool)
		}
		mapping[p[0]][p[1]] = true

		if _, ok := mapping[p[1]]; !ok {
			mapping[p[1]] = make(map[string]bool)
		}
		mapping[p[1]][p[0]] = true
	}

	len1 := len(words1)
	len2 := len(words2)

	if len1 != len2 {
		return false
	}

	for i := range words1 {
		if words1[i] == words2[i] {
			continue
		}

		if _, ok := mapping[words1[i]]; ok && mapping[words1[i]][words2[i]] {
			continue
		}

		return false
	}

	return true
}

//	problems
//	1.	I didn't aware that I assume each words exists only one
//		transformation, but this is not true
//	2.	optimization, the complexity is O(m * n * 2), m is length of words,
//		n is longest duplicates of a word mapping, and 2 is twice search
//		for reverse mapping.
//		bottleneck is at checking, if I can make search constant time, then
//		it will be much faster. The solution is to use a map
//		map[string]map[string]bool
