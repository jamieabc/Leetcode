package main

// Implement a SnapshotArray that supports the following interface:
//
//     SnapshotArray(int length) initializes an array-like data structure with the given length.  Initially, each element equals 0.
//     void set(index, val) sets the element at the given index to be equal to val.
//     int snap() takes a snapshot of the array and returns the snap_id: the total number of times we called snap() minus 1.
//     int get(index, snap_id) returns the value at the given index, at the time we took the snapshot with the given snap_id
//
//
//
// Example 1:
//
// Input: ["SnapshotArray","set","snap","set","get"]
// [[3],[0,5],[],[0,6],[0,0]]
// Output: [null,null,0,null,5]
// Explanation:
// SnapshotArray snapshotArr = new SnapshotArray(3); // set the length to be 3
// snapshotArr.set(0,5);  // Set array[0] = 5
// snapshotArr.snap();  // Take a snapshot, return snap_id = 0
// snapshotArr.set(0,6);
// snapshotArr.get(0,0);  // Get the value of array[0] with snap_id = 0, return 5
//
//
//
// Constraints:
//
//     1 <= length <= 50000
//     At most 50000 calls will be made to set, snap, and get.
//     0 <= index < length
//     0 <= snap_id < (the total number of times we call snap())
//     0 <= val <= 10^9

type Item struct {
	SnapID int
	Val    int
}

type SnapshotArray struct {
	Data   map[int][]Item
	SnapID int
}

func Constructor(length int) SnapshotArray {
	data := make(map[int][]Item)
	for i := 0; i < length; i++ {
		data[i] = make([]Item, 1)
		data[i][0] = Item{
			Val:    0,
			SnapID: 0,
		}
	}

	return SnapshotArray{
		Data:   data,
		SnapID: 0,
	}
}

func (this *SnapshotArray) Set(index int, val int) {
	item := this.Data[index]
	if item[len(item)-1].SnapID == this.SnapID {
		item[len(item)-1].Val = val
	} else {
		this.Data[index] = append(this.Data[index], Item{
			Val:    val,
			SnapID: this.SnapID,
		})
	}
}

func (this *SnapshotArray) Snap() int {
	this.SnapID++
	return this.SnapID - 1
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	item := this.Data[index]
	size := len(item)

	if item[size-1].SnapID <= snap_id {
		return item[size-1].Val
	}

	var i, j int
	for i, j = 0, size-2; i < j; {
		mid := i + (j-i+1)/2

		if item[mid].SnapID == snap_id {
			return item[mid].Val
		}

		if item[mid].SnapID < snap_id {
			i = mid
		} else {
			j = mid - 1
		}
	}

	return item[i].Val
}

/**
 * Your SnapshotArray object will be instantiated and called as such:
 * obj := Constructor(length);
 * obj.Set(index,val);
 * param_2 := obj.Snap();
 * param_3 := obj.Get(index,snap_id);
 */

//  problems
//  1.  too slow, every time snapshot, whole array is stored

//  2.  inspired from https://leetcode.com/problems/snapshot-array/discuss/350562/JavaPython-Binary-Search

//      store every index with version

//	3.	if an index is not set, need to fallback to find latest data

//	4.	if value exist, use that value

//	5.	cannot use binary search, can only search to first version exist

//	6.	from sample code, binary search is usable as long as store in
//		array with mono-increasing attributes

//	7.	when updating, if version ID equals, just update instead of
//		inserting

//	8. 	set tc: O(1)
//		get tc: O(log n), binary search for size n of specific index
