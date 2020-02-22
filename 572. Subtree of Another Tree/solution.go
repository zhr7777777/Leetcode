/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSubtree(s *TreeNode, t *TreeNode) bool {
    stack := []*TreeNode{}
    cur := s
    result := false
    for {
        if cur == nil && len(stack) == 0 {
            break
        }
        if cur != nil {
            stack = append(stack, cur)
            if(cur.Val == t.Val && isSameTree(cur, t)) {
                return true
            }
            cur = cur.Left
        } else {
            top := stack[len(stack)-1]
            cur = top.Right
            stack = stack[0:len(stack)-1]
        }
    }
    return result
}

func isSameTree(s *TreeNode, t *TreeNode) bool {
    if s == nil && t == nil {
        return true
    }
    if s == nil && t != nil {
        return false
    }
    if s != nil && t == nil {
        return false
    }
    if s.Val != t.Val {
        return false
    }
    return isSameTree(s.Left, t.Left) && isSameTree(s.Right, t.Right)
}