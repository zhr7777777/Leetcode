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
var reverseList = function(head) {
    if(!head) return head
    let headNode = new ListNode()
    headNode.next = null
    let pre = headNode
    let cur = head
    while(cur) {
        let next = cur.next
        cur.next = pre
        pre = cur
        cur = next
    }
    head.next = null
    return pre
};

// recursively
var reverseList = function(head) {
    if(!head) return head
    if(!head.next) return head
    let p = reverseList(head.next)
    head.next.next = head
    head.next = null
    return p
};