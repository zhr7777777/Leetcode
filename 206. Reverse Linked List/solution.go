/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
    pre := new(ListNode)
    pre = nil   // important pre := nil ( use of untyped nil )
    cur := head
    for {
        if cur == nil {
            break
        }
        next := cur.Next
        cur.Next = pre
        pre = cur
        cur = next
    }
    return pre
}

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    p := reverseList(head.Next)
    head.Next.Next = head
    head.Next =  nil
    return p
}