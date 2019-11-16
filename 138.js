// A linked list is given such that each node contains an additional random pointer which could point to any node in the list or null.
//
//     Return a deep copy of the list.
//
//     Example 1:
//
// Input:
// {"$id":"1","next":{"$id":"2","next":null,"random":{"$ref":"2"},"val":2},"random":{"$ref":"2"},"val":1}
//
// Explanation:
//     Node 1's value is 1, both of its next and random pointer points to Node 2.
// Node 2's value is 2, its next pointer points to null and its random pointer points to itself.
//
// Note:
//
//     You must return the copy of the given head as a reference to the cloned list.

/**
 * // Definition for a Node.
 * function Node(val,next,random) {
 *    this.val = val;
 *    this.next = next;
 *    this.random = random;
 * };
 */

function Node(val, next, random) {
    this.val = val;
    this.next = next;
    this.random = random;
}

/**
 * @param {Node} head
 * @return {Node}
 */
var copyRandomList = function (head) {
    if (head === null) {
        return null;
    }

    // duplicate
    let cur = head;
    while(cur !== null) {
        cur.next = new Node(cur.val, cur.next, null);
        cur = cur.next.next;
    }

    // deep copy
    cur = head;
    while(cur !== null) {
        cur.next.random = cur.random ? cur.random.next : null;
        cur = cur.next ? cur.next.next : null;
    }

    // destruct
    cur = head;
    let result = cur.next;
    let cur2 = result;
    while(cur !== null) {
        cur.next = cur.next ? cur.next.next : null;
        cur2.next = cur2.next ? cur2.next.next : null;

        cur = cur.next;
        cur2 = cur2.next;
    }

    return result;
};
