package main

// Implement the class ProductOfNumbers1 that supports two methods:
//
// 1. add(int num)
//
// Adds the number num to the back of the current list of numbers.
// 2. getProduct(int k)
//
// Returns the product of the last k numbers in the current list.
// You can assume that always the current list has at least k numbers.
// At any time, the product of any contiguous sequence of numbers will fit into a single 32-bit integer without overflowing.
//
//
//
// Example:
//
// Input
// ["ProductOfNumbers1","add","add","add","add","add","getProduct","getProduct","getProduct","add","getProduct"]
// [[],[3],[0],[2],[5],[4],[2],[3],[4],[8],[2]]
//
// Output
// [null,null,null,null,null,null,20,40,0,null,32]
//
// Explanation
// ProductOfNumbers1 productOfNumbers = new ProductOfNumbers1();
// productOfNumbers.add(3);        // [3]
// productOfNumbers.add(0);        // [3,0]
// productOfNumbers.add(2);        // [3,0,2]
// productOfNumbers.add(5);        // [3,0,2,5]
// productOfNumbers.add(4);        // [3,0,2,5,4]
// productOfNumbers.getProduct(2); // return 20. The product of the last 2 numbers is 5 * 4 = 20
// productOfNumbers.getProduct(3); // return 40. The product of the last 3 numbers is 2 * 5 * 4 = 40
// productOfNumbers.getProduct(4); // return 0. The product of the last 4 numbers is 0 * 2 * 5 * 4 = 0
// productOfNumbers.add(8);        // [3,0,2,5,4,8]
// productOfNumbers.getProduct(2); // return 32. The product of the last 2 numbers is 4 * 8 = 32
//
//
// Constraints:
//
// There will be at most 40000 operations considering both add and getProduct.
// 0 <= num <= 100
// 1 <= k <= 40000

type ProductOfNumbers struct {
	Product []int
	Prev    int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		Product: make([]int, 0),
		Prev:    1,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	if num != 0 {
		this.Prev *= num
		this.Product = append(this.Product, this.Prev)
	} else {
		this.Prev = 1
		this.Product = this.Product[:0]
	}
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	last := len(this.Product) - k

	if last >= 0 {
		if last > 0 {
			return this.Product[len(this.Product)-1] / this.Product[last-1]
		}
		return this.Product[len(this.Product)-1]
	}

	return 0
}

type ProductOfNumbers2 struct {
	Product []int
	Prev    int
	Zero    int
}

func Constructor2() ProductOfNumbers2 {
	return ProductOfNumbers2{
		Product: make([]int, 0),
		Prev:    1,
		Zero:    -1,
	}
}

func (this *ProductOfNumbers2) Add(num int) {
	if num != 0 {
		this.Prev *= num
	} else {
		this.Zero = len(this.Product)
		this.Prev = 1
	}
	this.Product = append(this.Product, this.Prev)
}

func (this *ProductOfNumbers2) GetProduct(k int) int {
	last := len(this.Product) - k

	if last > this.Zero {
		if last > 0 {
			return this.Product[len(this.Product)-1] / this.Product[last-1]
		}
		return this.Product[len(this.Product)-1]
	}

	return 0
}

type ProductOfNumbers1 struct {
	Product []int
	Zero    int
}

func Constructor1() ProductOfNumbers1 {
	return ProductOfNumbers1{
		Product: make([]int, 0),
		Zero:    -1,
	}
}

func (this *ProductOfNumbers1) Add(num int) {
	this.Product = append(this.Product, num)
}

func (this *ProductOfNumbers1) GetProduct(k int) int {
	last := len(this.Product) - k

	var product int

	if last > this.Zero {
		product = 1

		for i := len(this.Product) - 1; i >= last; i-- {
			product *= this.Product[i]
		}
	}

	return product
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */

//	Notes
//	1.	when needed, actual do multiplication

//	2.	store all product values is not necessary, because as long as there's
//		a zero, all previous product values will be cleared

//		when encounter a zero, reset array to length 0, it can also detect if
//		k > current product size means there's a zero somewhere in the range
