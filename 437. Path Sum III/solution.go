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
func pathSum(root *TreeNode, sum int) int {
    count := 0
    cur := new(TreeNodeWithVisited)
    cur.TreeNode = root
    cur.Visited = false
    stack := []*TreeNodeWithVisited{}
    candidates := []int{}
    for {
        if cur.TreeNode == nil && len(stack) == 0 {
            break
        }
        if cur.TreeNode != nil {
            for i:=0; i<len(candidates); i++ {
                if cur.TreeNode.Val == candidates[i] {
                    count++
                }
            }
            candidates = append(candidates, sum)
            for i:=0; i<len(candidates); i++ {
                candidates[i] = candidates[i] - cur.TreeNode.Val
            }
            
            stack = append(stack, cur)
            left := new(TreeNodeWithVisited)
            left.TreeNode = cur.TreeNode.Left
            left.Visited = false
            cur = left
        } else {
            top := new(TreeNodeWithVisited)
            top = stack[len(stack)-1]
            if top.Visited == true {
                if top.TreeNode.Val == sum {
                    count++
                }
                candidates = candidates[:len(candidates)-1]
                for i:=0; i<len(candidates); i++ {  // important should add back
                    candidates[i] = candidates[i] + top.TreeNode.Val
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
    return count
}