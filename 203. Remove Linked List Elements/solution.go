/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
    headNode := new(ListNode)   // important eg: [1, 2] 1
    headNode.Next = head
    cur := head
    pre := headNode
    for {
        if cur == nil {
            break
        }
        if cur.Val == val {
            pre.Next = cur.Next
            cur = cur.Next // [1] 1 is wrong
        } else {
            pre = cur
            cur = cur.Next
        }
    }
    return headNode.Next
}