package main

//Design a HashSet without using any built-in hash table libraries.
//
//To be specific, your design should include these functions:
//
//    add(value): Insert a value into the HashSet.
//    contains(value) : Return whether the value exists in the HashSet or not.
//    remove(value): Remove a value in the HashSet. If the value does not exist in the HashSet, do nothing.
//
//
//Example:
//
//MyHashSet hashSet = new MyHashSet();
//hashSet.add(1);
//hashSet.add(2);
//hashSet.contains(1);    // returns true
//hashSet.contains(3);    // returns false (not found)
//hashSet.add(2);
//hashSet.contains(2);    // returns true
//hashSet.remove(2);
//hashSet.contains(2);    // returns false (already removed)
//
//
//Note:
//
//    All values will be in the range of [0, 1000000].
//    The number of operations will be in the range of [1, 10000].
//    Please do not use the built-in HashSet library.

type Node struct {
	Val        int
	Next, Prev *Node
}

func (this *Node) Find(val int) *Node {
	if this.Val == val {
		return this
	}

	if this.Next == nil {
		return nil
	}

	return this.Next.Find(val)
}

func (this *Node) Add(val int) {
	if this.Val == val {
		return
	}

	if this.Next == nil {
		this.Next = &Node{
			Val:  val,
			Prev: this,
		}
		return
	}

	this.Next.Add(val)
}

type MyHashSet struct {
	Data []*Node
	Size int
}

/** Initialize your data structure here. */
func Constructor() MyHashSet {
	return MyHashSet{
		Data: make([]*Node, 100),
		Size: 100,
	}
}

func (this *MyHashSet) Add(key int) {
	bucket := key % this.Size

	if this.Data[bucket] == nil {
		// first one in the bucket, create
		this.Data[bucket] = &Node{
			Val: key,
		}
	} else {
		this.Data[bucket].Add(key)
	}
}

func (this *MyHashSet) Remove(key int) {
	bucket := key % this.Size

	if this.Data[bucket] == nil {
		return
	}

	n := this.Data[bucket].Find(key)

	if n == nil {
		return
	}

	if n == this.Data[bucket] {
		if n.Next != nil {
			this.Data[bucket] = n.Next
		} else {
			this.Data[bucket] = nil
		}
		return
	}

	if n.Prev != nil {
		n.Prev.Next = n.Next
	}

	if n.Next != nil {
		n.Next.Prev = n.Prev
	}
}

/** Returns true if this set contains the specified element */
func (this *MyHashSet) Contains(key int) bool {
	bucket := key % this.Size

	if this.Data[bucket] == nil {
		return false
	}

	if nil == this.Data[bucket].Find(key) {
		return false
	}
	return true
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */

//  problems
//  1.  be careful when found node is head node

//  2.  inspired from https://leetcode.com/problems/design-hashset/discuss/355141/Java-Solution

//      rehashing & BST
