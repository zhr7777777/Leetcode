/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNodeWithVisited struct {
    TreeNode *TreeNode
    Visited bool
}
func getPath(stack []*TreeNodeWithVisited) int {
    path := 0
    for i:=0; i<len(stack); i++ {
        path += stack[i].TreeNode.Val
    }
    return path
}
func isLeaf(node *TreeNode) bool {
    if node.Left == nil && node.Right == nil {
        return true
    }
    return false
}
func hasPathSum(root *TreeNode, sum int) bool {
    result := false
    cur := new(TreeNodeWithVisited)
    cur.TreeNode = root
    cur.Visited = false
    var stack []*TreeNodeWithVisited
    for {
        if cur.TreeNode == nil && len(stack) == 0 {
            break
        }
        if cur.TreeNode != nil {
            stack = append(stack, cur)
            left := new(TreeNodeWithVisited)
            left.TreeNode = cur.TreeNode.Left
            left.Visited = false
            cur = left
        } else {
            top := new(TreeNodeWithVisited)
            top = stack[len(stack)-1]
            if top.Visited == true {
                if isLeaf(top.TreeNode) == true && getPath(stack) == sum { // important, node must be a leaf
                    result = true
                    break
                }
                stack = stack[:len(stack)-1]
            } else {
                Right := new(TreeNodeWithVisited)
                Right.TreeNode = top.TreeNode.Right
                Right.Visited = false
                top.Visited = true
                cur = Right 
            }
        }
    }
    return result
}