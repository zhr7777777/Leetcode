/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return root
    }
    if root.Left == nil && root.Right == nil {
        return root
    }
    temp := root.Left
    root.Left = root.Right
    root.Right = temp
    root.Left = invertTree(root.Left)
    root.Right = invertTree(root.Right)
    return root
}