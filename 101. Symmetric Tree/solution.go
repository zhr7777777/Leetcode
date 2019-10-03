/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    return isSymmetricSameTree(root.Left, root.Right) 
}

func isSymmetricSameTree(p *TreeNode, q *TreeNode) bool {
    if p == nil && q == nil {   // important
        return true
    }
    if p == nil && q != nil {   // important
        return false
    }
    if p != nil && q == nil {   // important
        return false
    }
    return p.Val == q.Val && isSymmetricSameTree(p.Left, q.Right) && isSymmetricSameTree(p.Right, q.Left)
}