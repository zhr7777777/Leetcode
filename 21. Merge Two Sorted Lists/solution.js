/**
 * Definition for singly-linked list.
 * function ListNode(val) {
 *     this.val = val;
 *     this.next = null;
 * }
 */
/**
 * @param {ListNode} l1
 * @param {ListNode} l2
 * @return {ListNode}
 */
var mergeTwoLists = function(l1, l2) {
    var head1 = new ListNode(0)
    var cur1 = l1
    var cur2 = l2
    var pre1 = head1
    pre1.next = cur1
    while(cur1 && cur2) {
        if(cur2.val > cur1.val) {
            pre1 = cur1
            cur1 = cur1.next
        } else {
            let temp = cur2.next
            pre1.next = cur2
            cur2.next = cur1
            pre1 = cur2
            cur2 = temp
        }
    }
    if(cur2) {
        pre1.next = cur2
    }
    return head1.next
};