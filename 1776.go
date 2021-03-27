package main

// There are n cars traveling at different speeds in the same direction along a one-lane road. You are given an array cars of length n, where cars[i] = [positioni, speedi] represents:
//
// positioni is the distance between the ith car and the beginning of the road in meters. It is guaranteed that positioni < positioni+1.
// speedi is the initial speed of the ith car in meters per second.
//
// For simplicity, cars can be considered as points moving along the number line. Two cars collide when they occupy the same position. Once a car collides with another car, they unite and form a single car fleet. The cars in the formed fleet will have the same position and the same speed, which is the initial speed of the slowest car in the fleet.
//
// Return an array answer, where answer[i] is the time, in seconds, at which the ith car collides with the next car, or -1 if the car does not collide with the next car. Answers within 10-5 of the actual answers are accepted.
//
//
//
// Example 1:
//
// Input: cars = [[1,2],[2,1],[4,3],[7,2]]
// Output: [1.00000,-1.00000,3.00000,-1.00000]
// Explanation: After exactly one second, the first car will collide with the second car, and form a car fleet with speed 1 m/s. After exactly 3 seconds, the third car will collide with the fourth car, and form a car fleet with speed 2 m/s.
//
// Example 2:
//
// Input: cars = [[3,4],[5,4],[6,3],[9,1]]
// Output: [2.00000,1.00000,1.50000,-1.00000]
//
//
//
// Constraints:
//
// 1 <= cars.length <= 105
// 1 <= positioni, speedi <= 106
// positioni < positioni+1

func getCollisionTimes(cars [][]int) []float64 {
	size := len(cars)
	ans := make([]float64, size)
	for i := range ans {
		ans[i] = float64(-1)
	}

	stack := make([]int, 0)

	for i := size - 1; i >= 0; i-- {
		// current car catches next car after next car catches
		// next next car, which means next car is already becomes next next car
		// current car is slower or equal, never catches up
		for len(stack) > 0 &&
			(cars[i][1] <= cars[stack[len(stack)-1]][1] ||
				(len(stack) > 1 &&
					collideAt(cars[i], cars[stack[len(stack)-1]]) > ans[stack[len(stack)-1]])) {
			stack = stack[:len(stack)-1]
		}

		if len(stack) > 0 {
			ans[i] = collideAt(cars[i], cars[stack[len(stack)-1]])
		}

		stack = append(stack, i)
	}

	return ans
}

func collideAt(prev, next []int) float64 {
	if prev[1] == next[1] {
		return float64(-1)
	}
	return float64(next[0]-prev[0]) / float64(prev[1]-next[1])
}

//	Notes
//	1.	at first thought, should process by collide time, use a min-heap to store
//		each car collide time, select the one to collide closest, remove from
//		list, do it again

//		tc will be O(n^2 log(n)), in the worst case need each time only collide
//		one car, need n rounds, and each time need to put remaining n cars into
//		heap, each push operation takes log(n)

//		this is not efficient, so need to change to another way of thinking,
//		since faster car will be block by slower car, and faster car is vanished,
//		it only cares about which car behind catches up, and this cur collide
//		to next car.

//		it seems like stack operation, the question becomes how to make sure
//		greedy operation guarantee to find answer

//		the reason is because slower car blocks faster car, as long as previous
//		car collide to current car is earlier than current car collide to next
//		car, result is true

//		if current car collides to next car is earlier than prev car collide
//		to current car, move on to next car, because
//		speed(prev) > speed(cur) > speed(next), so as long as next car is
//		processed, all conditions will be considered

//	2.	inspired from https://leetcode.com/problems/car-fleet-ii/discuss/1085987/JavaC%2B%2BPython-O(n)-Stack-Solution

//		slower car blocks faster car, if faster -> slower, faster collided
//		and vanish, check previous car earlier than faster

//		because collision only happens when faster -> slower, slower car will
//		never be affected, thus start from last car and find any previous car
//		which is faster

//	3.	inspired from https://leetcode.com/problems/car-fleet-ii/discuss/1086115/C%2B%2B-Stack-minimalizm

//		a -> b -> c -> d
//		stack: [d]

//		if c.speed >= d.speed, put into stack, it will be considered later,
//		also, collision time of c can be calculated

//		stack: [d, c]
//		- if b.speed <= d.speed (slowest), stack is cleared because b blocks
//		  all cars after (in this case is a)

//		- if c catches d before b catches c, which means c is already vanished,
//		  thus c can be removed

//	4.	inspired from https://leetcode.com/problems/car-fleet-ii/discuss/1086204/JavaScript-Stack-Solution-w-explanation

//		I can understand only after reading this post

// 		There are a few things that we know (in this example I will use cars
// 		[a, b, c, d, e]):
//
// 		- a can hit b, c, d, and e
// 		- a can either hit b before b hits any other car or after
// 		- If a hits b before b hits any other car, then the calculation is the
// 	      distance from a to b divided by the difference in speeds
// 		- If a hits b after b hits c, a is really hitting c because once b
// 		  collides with c b becomes c's speed
// 		- Utilizing the knowledge from the bullet above, if a hits b or c after
// 		  b or c hits d, a is really hitting d
// 		- If there are no cars available to be hit then the result for that car
// 		  is -1.
