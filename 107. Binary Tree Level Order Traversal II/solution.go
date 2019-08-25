/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrderBottom(root *TreeNode) [][]int {
    var result [][]int
    if root == nil {
        return result
    }
    queue := []*TreeNode{root}
    for {
        if len(queue) == 0 {
            break
        }
        var curLevel []int
        var nextQueue []*TreeNode
        for i:=0; i<len(queue); i++ {
            curLevel = append(curLevel, queue[i].Val)
            if queue[i].Left != nil {
                nextQueue = append(nextQueue, queue[i].Left)
            }
            if queue[i].Right != nil {
                nextQueue = append(nextQueue, queue[i].Right)
            }
        }
        result = append([][]int{curLevel}, result...)
        queue = nextQueue
    }
    return result
}