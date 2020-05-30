/**
 * Definition for singly-linked list.
 * function ListNode(val, next) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.next = (next===undefined ? null : next)
 * }
 */
/**
 * @param {ListNode} head
 * @return {boolean}
 */
var isPalindrome = function(head) {
    if(!head || !head.next) return true
    if(!head.next.next && head.val !== head.next.val) {
        return false
    }
    if(!head.next.next && head.val === head.next.val) {
        return true
    }
    let half = []
    let s1 = head
    let s2 = head
    while(s2 && s2.next) {
        half.push(s1.val)
        s1 = s1.next
        s2 = s2.next.next
    }
    if(s2) {
        s1 = s1.next
    }
    for(let i=0; i<half.length; i++) {
        if(half[half.length-1-i] !== s1.val) {
            return false
        }
        s1 = s1.next
    }
    return true
};