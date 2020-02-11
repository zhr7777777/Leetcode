/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var tilt int = 0
func findTilt(root *TreeNode) int {
    tilt = 0
    dfs(root)
    return tilt
}

func dfs(root *TreeNode) int {
    if root == nil {
        return 0
    }
    left := dfs(root.Left)
    right := dfs(root.Right)
    tilt += abs(left - right)
    return root.Val + left + right
}

func abs(n int) int {
    return int(math.Abs(float64(n)))
}