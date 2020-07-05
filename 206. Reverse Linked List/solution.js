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

// Input: 1->2->3->4->5->NULL
// Output: 5->4->3->2->1->NULL

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
    if(!head) { // 考虑空链表的情况
        return null
    }
    if(!head.next) return head
    let p = reverseList(head.next)  // 递归到最后一个节点，比如上面的5
    head.next.next = head   // 对于非最后节点的操作
    head.next = null
    return p
};