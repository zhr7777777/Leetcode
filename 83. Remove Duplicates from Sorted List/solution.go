/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
    var firstNode ListNode
    var first *ListNode = &firstNode
    var pre *ListNode = head
    var cur *ListNode = head
    first.Next = head
    for {
        if cur == nil { // important
            break
        }
        if cur.Next == nil {
            if pre != cur { // important
                pre.Next = nil
            }
            break
        }
        if cur.Val == cur.Next.Val {
            cur = cur.Next
        } else {
            pre.Next = cur.Next
            pre = cur.Next
            cur = cur.Next
        }
    }
    return first.Next
}