/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var middleNode = function(head) {
    if(!head.next) return head
    if(!head.next.next) return head.next
    let s1 = head
    let s2 = head
    while(s2 && s2.next) {
        s1 = s1.next
        s2 = s2.next.next
    }
    return s1
};