/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    curMax := 1
    depth(root, &curMax)
    return curMax - 1
}

func depth(root *TreeNode, curMax *int) int {
    if root == nil {
        return 0
    }
    l := depth(root.Left, curMax)
    r := depth(root.Right, curMax)
    *curMax = int(math.Max(float64(*curMax), float64(l+r+1)))
    return 1 + int(math.Max(float64(l), float64(r)))
}