/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {
        return l2
    }
    if l2 ==nil {
        return l1
    }
    var headNode ListNode
    headNode.Val = 0
    headNode.Next = nil
    var head *ListNode = &headNode
    var cur *ListNode
    if l1.Val >= l2.Val {
        head.Next = l2
        l2 = l2.Next
    } else {
        head.Next = l1
        l1 = l1.Next
    }
    cur = head.Next
    for {
        if l1 == nil {
            cur.Next = l2
            break
        }
        if l2 == nil {
            cur.Next = l1
            break
        }
        if l1.Val >= l2.Val {
            cur.Next = l2
            cur = l2
            l2 = l2.Next
        } else {
            cur.Next = l1
            cur = l1
            l1 = l1.Next
        }
    }
    return head.Next
}