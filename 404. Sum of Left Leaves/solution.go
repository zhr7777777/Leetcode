/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumOfLeftLeaves(root *TreeNode) int {
    stack := []*TreeNode{}
    cur := root
    sum := 0
    for {
        if cur == nil && len(stack) == 0 {
            break
        }
        if cur != nil {
            stack = append(stack, cur)
            cur = cur.Left
            if cur != nil && cur.Left == nil && cur.Right == nil {
                sum += cur.Val
            }
        } else {
            top := stack[len(stack)-1]
            cur = top.Right
            stack = stack[:len(stack)-1]
        }
    }
    return sum
}