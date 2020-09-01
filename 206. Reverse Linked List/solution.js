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

const reverseList = node => {
  let head = new ListNode()
  head.next = node
  let cur = node
  let pre = head
  while(cur) {
    let temp = cur.next
    cur.next = pre
    pre = cur
    cur = temp
  }
  // head = null // 这里把head置位null无效，因为只是把head指向null，node.next还是指向原来head指向的head node
  node.next = null
  return pre
}

// recursive
const reverseList = node => {
    if(node === null) return null // 无node情况
    if(node.next === null) return node  // 一个node情况
    const lastNode = reverseList(node.next)
    node.next.next = node // 两个node情况
    node.next = null
    return lastNode
}