/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} head
 * @return {boolean}
 */
var hasCycle = function(head) {
    let cur = head
    let records = new Map()
    while(cur) {
        if(records.has(cur)) {
            return true
        }
        records.set(cur, 1)
        cur = cur.next
    }
    return false
};