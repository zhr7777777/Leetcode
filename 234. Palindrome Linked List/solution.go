/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    if head == nil || head.Next == nil {
        return true
    }
    if head.Next.Next == nil {
        if head.Val == head.Next.Val {
            return true
        } else {
            return false
        }
    }
    step1 := head
    step2 := head
    for {
        step1 = step1.Next
        step2  = step2.Next.Next
        if step2.Next == nil || step2.Next.Next == nil {
            break
        }
    }
    cur1 := head
    cur2 := reverseLinkedList(step1.Next)
    for {
        if cur1 == nil || cur2 == nil {
            break
        }
        if cur1.Val != cur2.Val {
            return false
        }
        cur1 = cur1.Next
        cur2 = cur2.Next
    }
    return true
}

func reverseLinkedList(head *ListNode) *ListNode {
    if head.Next == nil {
        return head
    }
    p := reverseLinkedList(head.Next)
    head.Next.Next = head
    head.Next = nil
    return p
}