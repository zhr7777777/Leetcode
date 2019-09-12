/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
    var curA *ListNode = headA
    var curB *ListNode = headB
    var intersectedNode *ListNode
    if curA == nil || curB == nil {
        return nil
    } 
    flagA := false
    flagB := false
    for {
        if curA == curB {
            intersectedNode = curA
            break
        }
        if flagA && flagB && curA == nil && curB == nil {
            break
        }
        if curA == nil {
            curA = headB
            flagA = true
            continue    // important
        }
        if curB == nil {
            curB = headA
            flagB = true
            continue    // important
        }
        curA = curA.Next
        curB = curB.Next
    }
    return intersectedNode
}