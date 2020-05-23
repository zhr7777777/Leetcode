/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @param {number} k
 * @return {boolean}
 */
var findTarget = function(root, k) {
    let map = {}
    let stack = []
    let cur = root
    while(cur || stack.length) {
        if(cur) {
            stack.push(cur)
            cur = cur.left
        } else {
            let top = stack.pop()
            let target = k - top.val
            if(target in map) {
                return true
            } else {
                map[top.val] = top.val
            }
            cur = top.right
        }
    }
    return false
};