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
var deleteDuplicates = function(head) {
    let headNode = new ListNode(Symbol())
    headNode.next = head
    let pre = headNode
    let cur = head
    while(cur) {
        if(cur.val === pre.val) {
            pre.next = pre.next.next
            cur = pre.next
        } else {
            pre = cur
            cur = cur.next
        }
    }
    return headNode.next
};