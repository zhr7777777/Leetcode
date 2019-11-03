// type ListNode struct {
//     Val int
//     Next *ListNode
// }

// 276ms
// func isSubsequence(s string, t string) bool {
//     var mappingT map[byte]*ListNode
//     if mappingT == nil {
//         mappingT = make(map[byte]*ListNode)
//         for i:=0; i<len(t); i++ {
//             headNode, has := mappingT[t[i]]
//             cur := new(ListNode)
//             cur.Val = i
//             mappingT[t[i]] = cur
//             if has {
//                 cur.Next = headNode
//             }
//         }    
//     }
//     previousIndex := -1
//     for i:=0; i<len(s); i++ {
//         cur, has := mappingT[s[i]]
//         if !has {
//             return false
//         }
//         for {
//             if cur.Val >= previousIndex {
//                 if cur.Next != nil && cur.Next.Val >= previousIndex {    // cur.Next.Val >= previousIndex is important
//                     cur = cur.Next
//                     continue
//                 } else {
//                     previousIndex = cur.Val
//                     cur.Val = -1 // important
//                     break   
//                 }
//             }
//             if cur.Next == nil {
//                 return false
//             }
//             cur = cur.Next
//         }
//     }
//     return true
// }

// 12ms
func isSubsequence(s string, t string) bool {
    if s == "" {
        return true
    }
    index := 0
    for i:=0; i<len(t); i++ {
        if t[i] == s[index] {
            if index == len(s) - 1 {    // important
                return true
            }
            index++
        }
    }
    if index == len(s) { // not len(s) - 1
        return true
    }
    return false
}