package main

//Design and implement a data structure for a compressed string iterator. It should support the following operations: next and hasNext.
//
//The given compressed string will be in the form of each letter followed by a positive integer representing the number of this letter existing in the original uncompressed string.
//
//next() - if the original string still has uncompressed characters, return the next letter; Otherwise return a white space.
//hasNext() - Judge whether there is any letter needs to be uncompressed.
//
//Note:
//Please remember to RESET your class variables declared in StringIterator, as static/class variables are persisted across multiple test cases. Please see here for more details.
//
//Example:
//
//StringIterator iterator = new StringIterator("L1e2t1C1o1d1e1");
//
//iterator.next(); // return 'L'
//iterator.next(); // return 'e'
//iterator.next(); // return 'e'
//iterator.next(); // return 't'
//iterator.next(); // return 'C'
//iterator.next(); // return 'o'
//iterator.next(); // return 'd'
//iterator.hasNext(); // return true
//iterator.next(); // return 'e'
//iterator.hasNext(); // return false
//iterator.next(); // return ' '

type StringIterator struct {
	str           string
	char          byte
	remainCount   int
	nextCharIndex int
}

func Constructor(compressedString string) StringIterator {
	s := StringIterator{
		str:  compressedString,
		char: ' ',
	}

	s.remainCount, s.nextCharIndex = parse(s.str, 0)

	if len(s.str) > 0 {
		s.char = s.str[0]
	}

	return s
}

// index - location of character, so number starts from next one
func parse(str string, index int) (int, int) {
	var end, total int
	length := len(str)

	for end = index + 1; end < length; end++ {
		if isNum(str[end]) {
			total *= 10
			total += int(str[end] - '0')
		} else {
			break
		}
	}

	return total, end
}

func isNum(c byte) bool {
	return c-'0' >= 0 && c-'0' <= 9
}

func (this *StringIterator) Next() byte {
	if !this.HasNext() {
		return ' '
	}

	if this.remainCount > 0 {
		this.remainCount--
		return this.char
	}

	this.char = this.str[this.nextCharIndex]
	this.remainCount, this.nextCharIndex = parse(this.str, this.nextCharIndex)
	this.remainCount--

	return this.char
}

func (this *StringIterator) HasNext() bool {
	if this.remainCount == 0 {
		return this.nextCharIndex != len(this.str)
	}

	return true
}

/**
 * Your StringIterator object will be instantiated and called as such:
 * obj := Constructor(compressedString);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */

//	problems
//	1.	number could be multiple, I cannot assume number is always 1 digit
//	2.	when no next, return space instead of \u0000
//	3. 	instead of using strconv.Atoi, can I do it myself to faster
//		calculation?
//	4.	optimize, parse whole string at the beginning, then for all other
//		operations are just query
//	5.	optimize, parse all chars at beginning may be a waste if only part
//		of chars are used. It should be checking on the fly.
//	6.	optimize, store char instead of index, reduce additional get char
//	7.	optimize, targetCount can be used to check, if it's 0
//	8.	optimize, remainCount no need additional minus, just change checking
//		criteria to 1
