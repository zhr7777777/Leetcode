/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func minDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    treeNodeQueue := []*TreeNode{root}
    minDepth := 1
    for {
        if len(treeNodeQueue) == 0 {
            break
        }
        var treeLayerNodes []*TreeNode
        for i:=0; i<len(treeNodeQueue); i++ {
            if treeNodeQueue[i].Left == nil && treeNodeQueue[i].Right == nil {
                return minDepth
                break
            }
            if treeNodeQueue[i].Left != nil {
                treeLayerNodes = append(treeLayerNodes, treeNodeQueue[i].Left)    
            }
            if treeNodeQueue[i].Right != nil {
                treeLayerNodes = append(treeLayerNodes, treeNodeQueue[i].Right)    
            }
        }
        treeNodeQueue = treeLayerNodes
        minDepth++
    }
    return minDepth
}