/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */

/**
 * @param {ListNode} headA
 * @param {ListNode} headB
 * @return {ListNode}
 */
var getIntersectionNode = function(headA, headB) {
    let cur1 = headA
    let cur2 = headB
    let f1 = false
    let f2 = false
    while(cur1 && cur2) {
        if(cur1 === cur2) {
            return cur1
        }
        cur1 = cur1.next
        cur2 = cur2.next
        if(cur1 === null && !f1) {
            f1 = true
            cur1 = headB
        }
        if(cur2 === null && !f2) {
            f2 = true
            cur2 = headA
        }
    }
    return null
};