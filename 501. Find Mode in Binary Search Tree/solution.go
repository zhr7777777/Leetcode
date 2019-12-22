/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findMode(root *TreeNode) []int {
    modes := []int{}
    curMaxCount := 0
    stack := []*TreeNode{}
    cur := root
    for {
        if cur == nil && len(stack) == 0 {
            break
        }
        if cur != nil {
            stack = append(stack, cur)
            if len(modes) == 0 {
                modes = append(modes, cur.Val)
                curMaxCount = countSameLeaf(cur, cur)
            } else {
                if cur.Val == modes[len(modes)-1] {
                    
                } else {
                    curCount := countSameLeaf(cur, cur)
                    if curCount > curMaxCount {
                        curMaxCount = curCount
                        modes = []int{}
                        modes = append(modes, cur.Val)
                    } else if curCount == curMaxCount {
                        modes = append(modes, cur.Val)
                    }
                }
            }
            cur = cur.Left  
        } else {
            top := stack[len(stack)-1]
            cur = top.Right
            stack = stack[:len(stack)-1]
        }
        
    }
    return modes
}


func countSameLeaf(leaf *TreeNode, root *TreeNode) int {
    if leaf == nil {
        return 0
    }
    if leaf.Val < root.Val {
        return countSameLeaf(leaf.Right, root)
    }
    if leaf.Val > root.Val {
        return countSameLeaf(leaf.Left, root)
    }
    return countSameLeaf(leaf.Left, root) + countSameLeaf(leaf.Right, root) + 1
}