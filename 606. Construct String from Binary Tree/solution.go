/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func tree2str(t *TreeNode) string {
    if t == nil {
        return ""
    }
    str := postTraverseAndGeneStr(t)
    return str[1:len(str)-1]
}

func postTraverseAndGeneStr(t *TreeNode) string {
    if t == nil {
        return ""
    }
    if t != nil && t.Left == nil && t.Right != nil {
        return "(" + strconv.Itoa(t.Val) + "()" + postTraverseAndGeneStr(t.Right) + ")"
    }
    return "(" + strconv.Itoa(t.Val) + postTraverseAndGeneStr(t.Left) + postTraverseAndGeneStr(t.Right) + ")"
}