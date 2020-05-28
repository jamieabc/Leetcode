package main

import "strings"

// Given a list of directory info including directory path, and all the files with contents in this directory, you need to find out all the groups of duplicate files in the file system in terms of their paths.
//
// A group of duplicate files consists of at least two files that have exactly the same content.
//
// A single directory info string in the input list has the following format:
//
// "root/d1/d2/.../dm f1.txt(f1_content) f2.txt(f2_content) ... fn.txt(fn_content)"
//
// It means there are n files (f1.txt, f2.txt ... fn.txt with content f1_content, f2_content ... fn_content, respectively) in directory root/d1/d2/.../dm. Note that n >= 1 and m >= 0. If m = 0, it means the directory is just the root directory.
//
// The output is a list of group of duplicate file paths. For each group, it contains all the file paths of the files that have the same content. A file path is a string that has the following format:
//
// "directory_path/file_name.txt"
//
// Example 1:
//
// Input:
// ["root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"]
// Output:
// [["root/a/2.txt","root/c/d/4.txt","root/4.txt"],["root/a/1.txt","root/c/3.txt"]]
//
//
//
// Note:
//
//     No order is required for the final output.
//     You may assume the directory name, file name and file content only has letters and digits, and the length of file content is in the range of [1,50].
//     The number of files given is in the range of [1,20000].
//     You may assume no files or directories share the same name in the same directory.
//     You may assume each given directory info represents a unique directory. Directory path and file info are separated by a single blank space.
//
//
// Follow-up beyond contest:
//
//     Imagine you are given a real file system, how will you search files? DFS or BFS?
//     If the file content is very large (GB level), how will you modify your solution?
//     If you can only read the file by 1kb each time, how will you modify your solution?
//     What is the time complexity of your modified solution? What is the most time-consuming part and memory consuming part of it? How to optimize?
//     How to make sure the duplicated files you find are not false positive?

func findDuplicate(paths []string) [][]string {
	mapping := make(map[string][]string)

	for _, p := range paths {
		strs := strings.Split(p, " ")
		dir := strs[0]
		for i := 1; i < len(strs); i++ {
			filename, content := extract(strs[i])
			mapping[content] = append(mapping[content], dir+"/"+filename)
		}
	}

	result := make([][]string, 0)
	for _, v := range mapping {
		if len(v) > 1 {
			result = append(result, v)
		}
	}

	return result
}

func extract(str string) (string, string) {
	for i := len(str) - 1; i >= 0; i-- {
		if str[i] == '(' {
			return str[:i], str[i+1 : len(str)-1]
		}
	}

	return "", ""
}

//	problems
//	1.	target is to find duplicates, not find group

//	2.	add reference https://leetcode.com/problems/find-duplicate-file-in-system/discuss/104123/C%2B%2B-clean-solution-answers-to-follow-up

//		DFS is more reasonable, since it's more easily to have 100 files
//		in a directory, instead of having 100 sub-directories, BFS might
//		in general w/ more memory usage

//		also, it's a good way to compare meta data first (e.g. file size),
//		then hash file with a small part of content, at last is byte by
//		byte comparison

//		time complexity is O(n^2) cause it might need to compare all others

//	3.	add reference https://leetcode.com/problems/find-duplicate-file-in-system/discuss/104120/Follow-up-questions-discussion
