/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }
    root := new(TreeNode)
    root.Val = nums[int(len(nums)/2)]
    root.Left = sortedArrayToBST(nums[0:int(len(nums)/2)])
    root.Right = sortedArrayToBST(nums[int(len(nums)/2)+1:len(nums)])
    return root
}