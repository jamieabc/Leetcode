package main

//Design a class to find the kth largest element in a stream. Note that it is the kth largest element in the sorted order, not the kth distinct element.
//
//Your KthLargest class will have a constructor which accepts an integer k and an integer array nums, which contains initial elements from the stream. For each call to the method KthLargest.add, return the element representing the kth largest element in the stream.
//
//Example:
//
//int k = 3;
//int[] arr = [4,5,8,2];
//KthLargest kthLargest = new KthLargest(3, arr);
//kthLargest.add(3);   // returns 4
//kthLargest.add(5);   // returns 5
//kthLargest.add(10);  // returns 5
//kthLargest.add(9);   // returns 8
//kthLargest.add(4);   // returns 8
//Note:
//You may assume that nums' length ≥ k-1 and k ≥ 1.

type minPQ struct {
	data   []int // data length should be k+1, make sure data never filled full
	length int
}

func (p minPQ) size() int {
	return p.length
}

func (p minPQ) peek() int {
	return p.data[0]
}

// borrow from java, add an element into pq
func (p *minPQ) offer(i int) {
	p.data[p.length] = i
	p.length++

	idx := bubbleUp(p.data, p.length-1)
	bubbleDown(p.data, idx, p.length)
}

func parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

func leftChild(i, size int) int {
	left := i*2 + 1
	if left >= size {
		return -1
	}
	return left
}

func rightChild(i, size int) int {
	right := i*2 + 2
	if right >= size {
		return -1
	}
	return right
}

func bubbleUp(data []int, i int) int {
	p := parent(i)
	for i > 0 {
		if data[i] >= data[p] {
			return i
		}
		swap(data, i, p)
		i = p
		p = parent(p)
	}
	return i
}

func swap(data []int, src, dst int) {
	data[src], data[dst] = data[dst], data[src]
}

// remove an element from pq
func (p *minPQ) poll() int {
	popped := p.data[0]
	p.data[0] = p.data[p.length-1]
	p.length--

	bubbleDown(p.data, 0, p.length)

	return popped
}

func bubbleDown(data []int, i, length int) {
	for {
		l := leftChild(i, length)
		r := rightChild(i, length)

		if l < 0 && r < 0 {
			return
		}

		// both exist
		if l > 0 && r > 0 {
			if data[l] >= data[i] && data[r] >= data[i] {
				return
			}

			if data[l] < data[r] {
				if data[l] >= data[i] {
					return
				}
				swap(data, i, l)
				i = l
			} else {
				if data[r] >= data[i] {
					return
				}
				swap(data, i, r)
				i = r
			}
			continue
		}

		// only one child exists
		if l < 0 {
			if data[r] >= data[i] {
				return
			}
			swap(data, i, r)
			i = r
		} else {
			if data[l] >= data[i] {
				return
			}
			swap(data, i, l)
			i = l
		}
	}
}

type KthLargest struct {
	pq *minPQ
	k  int
}

func Constructor(k int, nums []int) KthLargest {
	kth := KthLargest{
		pq: &minPQ{
			length: 0,
			data:   make([]int, k+1),
		},
		k: k,
	}

	// maintain max kth data
	for i := 0; i < len(nums); i++ {
		kth.pq.offer(nums[i])
		if kth.pq.size() > k {
			_ = kth.pq.poll()
		}
	}

	return kth
}

func (this *KthLargest) Add(val int) int {
	this.pq.offer(val)
	if this.pq.size() > this.k {
		_ = this.pq.poll()
	}
	return this.pq.peek()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */

// problems
// 1. sort.Ints in ascending order, but needs to be descending order
// 2. when second copy, should starts from i+1 instead of whole array (starts from 0)
// 3. not always increase element, only if some number is inserted
// 4. not always increase element, so this.data should not always update
// 5. this.data & this.length only updates when something is inserted
// 6. if first element is the case, then length is not updated
// 7. if first element updated, original this.data is not updated
// 8. length is updated only if element added
// 9. when length < k, additional element should be added directly
// 10. use priority queue to accelerate speed
