package main

//Given a function rand7 which generates a uniform random integer in the range 1 to 7, write a function rand10 which generates a uniform random integer in the range 1 to 10.
//
//Do NOT use system's Math.random().
//
//
//
//Example 1:
//
//Input: 1
//Output: [7]
//
//Example 2:
//
//Input: 2
//Output: [8,4]
//
//Example 3:
//
//Input: 3
//Output: [8,1,10]
//
//
//
//Note:
//
//    rand7 is predefined.
//    Each testcase has one argument: n, the number of times that rand10 is called.
//
//
//
//Follow up:
//
//    What is the expected value for the number of calls to rand7() function?
//    Could you minimize the number of calls to rand7()?

func rand10() int {
	tmp := 49
	for tmp > 40 {
		tmp = 7*(rand7()-1) + rand7()
	}

	if tmp%10 == 0 {
		return 10
	}
	return tmp % 10
}

//	Notes
//	1.	didn't think of solution in the first time

//	2.	inspired from https://leetcode.com/problems/implement-rand10-using-rand7/discuss/150301/Three-line-Java-solution-the-idea-can-be-generalized-to-%22Implement-RandM()-Using-RandN()%22

//		main idea: 1-7 is evenly distributed, 1-49 is also evenly distributed.

//		1-49 discard 41-49 becomes 1-40, 1-40 % 10 => 10

//	3.	inspired from https://leetcode.com/problems/implement-rand10-using-rand7/discuss/338395/In-depth-straightforward-detailed-explanation.-(Short-Java-solution)

//		author provides clear thinking process about random number generation
