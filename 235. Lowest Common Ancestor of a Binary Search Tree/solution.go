/**
 * Definition for TreeNode.
 * type TreeNode struct {
 *     Val int
 *     Left *ListNode
 *     Right *ListNode
 * }
 */
type TreeNodeWithVisited struct {
    TreeNode *TreeNode
    Visited bool
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    pParents := findTreeNodePath(root, p)
    qParents := findTreeNodePath(root, q)
    arr := []*TreeNode{}
    for i:=0; i<min(len(pParents), len(qParents)); i++ {
        if pParents[i].Val == qParents[i].Val {
            arr = append(arr, pParents[i])
        }
    }
    return arr[len(arr)-1]
}

func min(p, q int) int {
    if q <= p {
        return q
    }
    return p
}

func findTreeNodePath(root, p *TreeNode) []*TreeNode {
    arr := []*TreeNode{}
    cur := root
    for {
        if p.Val == cur.Val {
            arr = append(arr, cur)
            break
        }
        if p.Val > cur.Val {
            arr = append(arr, cur)
            cur = cur.Right
        } else if p.Val < cur.Val {
            arr = append(arr, cur)
            cur = cur.Left
        }
    }
    return arr
}


// func findTreeNodePath(root, p *TreeNode) []*TreeNode {
//     cur := new(TreeNodeWithVisited)
//     cur.TreeNode = root
//     cur.Visited = false
//     stack := []*TreeNodeWithVisited{}
//     result := []*TreeNode{}
//     for {
//         if cur.TreeNode == nil && len(stack) == 0 {
//             break
//         }
//         if cur.TreeNode != nil {
//             stack = append(stack, cur)
//             left := new(TreeNodeWithVisited)
//             left.TreeNode = cur.TreeNode.Left
//             left.Visited = false
//             cur = left
//         } else {
//             top := stack[len(stack)-1]
//             if !top.Visited {
//                 top.Visited = true
//                 right := new(TreeNodeWithVisited)
//                 right.TreeNode = top.TreeNode.Right
//                 right.Visited = false
//                 cur = right 
//             } else {
//                 if top.TreeNode.Val == p.Val {
//                     break
//                 }
//                 stack = stack[:len(stack)-1]
//             }
//         }
//     }
//     for i:=0; i<len(stack); i++ {
//         result = append(result, stack[i].TreeNode)
//     }
//     return result
// }