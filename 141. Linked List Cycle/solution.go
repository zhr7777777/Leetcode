/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func hasCycle(head *ListNode) bool {
    if head == nil {
        return false
    }
    result := false
    cur := new(ListNode)
    cur = head
    for {
        if cur.Next == nil {
            break
        }
        if cur.Val == 10000 {
            result = true
            break
        }
        cur.Val = 10000
        cur = cur.Next
    }
    return result
}