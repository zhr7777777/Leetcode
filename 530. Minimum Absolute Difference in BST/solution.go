/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func getMinimumDifference(root *TreeNode) int {
    miniDiff := getMiniDiff(root)
    stack := []*TreeNode{}
    cur := root
    for {
        if (cur == nil && len(stack) == 0) || miniDiff == 0 {
            break
        }
        if cur != nil {
            stack = append(stack, cur)
            if cur != root && (cur.Left != nil || cur.Right != nil) {
                curDiff := getMiniDiff(cur)
                if curDiff < miniDiff {
                    miniDiff = curDiff
                }
            }
            cur = cur.Left
        } else {
            top := stack[len(stack)-1]
            cur = top.Right
            stack = stack[:len(stack)-1]
        }
    }
    return miniDiff
}

func getLeftMax(left *TreeNode) *TreeNode {
    if left == nil {
        return nil
    }
    for {
        if left.Right == nil {  // important
            return left
        }
        left = left.Right
    }
    return nil
}

func getRightMin(right *TreeNode) *TreeNode {
    if right == nil {
        return nil
    }
    for {
        if right.Left == nil {  // important
            return right
        }
        right = right.Left
    }
    return nil
}

func getDis(a int, b int) int {
    return int(math.Abs(float64(a-b)))
}

func getMiniDiff(root *TreeNode) int {
    miniDiff := 0
    leftMaxNode := getLeftMax(root.Left)
    rightMinNode := getRightMin(root.Right)
    if leftMaxNode == nil {
        miniDiff = getDis(root.Val, rightMinNode.Val)
    } else if rightMinNode == nil {
        miniDiff = getDis(root.Val, leftMaxNode.Val)
    } else {
        miniDiff = getDis(root.Val, rightMinNode.Val)
        leftDiff := getDis(root.Val, leftMaxNode.Val)
        if leftDiff < miniDiff {
            miniDiff = leftDiff
        }
    }
    return miniDiff
}

