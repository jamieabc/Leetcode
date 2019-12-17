package main

//Count the number of segments in a string, where a segment is defined to be a contiguous sequence of non-space characters.
//
//Please note that the string does not contain any non-printable characters.
//
//Example:
//
//Input: "Hello, my name is John"
//Output: 5

func countSegments(s string) int {
	length := len(s)

	if length == 0 {
		return 0
	}

	var j int
	count := 0

	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			continue
		}

		count++
		for j = i + 1; j < length; j++ {
			if s[j] == ' ' {
				break
			}
		}

		i = j
	}

	return count
}
