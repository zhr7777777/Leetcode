/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    diff := maxDepth(root.Left) - maxDepth(root.Right)
    return diff >= -1 && diff <=1 && isBalanced(root.Left) && isBalanced(root.Right) // subtrees should also be balanced
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}

func max(a int, b int) int {
    if a >= b {
        return a
    } else {
        return b
    }
}