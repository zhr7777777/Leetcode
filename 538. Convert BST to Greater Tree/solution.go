/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func convertBST(root *TreeNode) *TreeNode {
    stack := []*TreeNode{}
    cur := root
    sum := 0
    for {
        if len(stack) == 0 && cur == nil {
            break
        }
        if cur != nil {
            stack = append(stack, cur)
            cur = cur.Right
        } else {
            top := stack[len(stack)-1]
            sum += top.Val
            top.Val = sum
            cur = top.Left
            stack = stack[0:len(stack)-1]
        }
    }
    return root
}