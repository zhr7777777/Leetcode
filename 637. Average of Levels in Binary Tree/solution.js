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
 * @return {number[]}
 */
var averageOfLevels = function(root) {
    let nodes = [root]
    let result = [root.val]
    while(nodes.length) {
        let sum = 0
        let nums = 0
        let length = nodes.length
        for(let i=0; i<length; i++) {
            if(nodes[i].left) {
                nodes.push(nodes[i].left)
                sum += nodes[i].left.val
                nums++
            }
            if(nodes[i].right) {
                nodes.push(nodes[i].right)
                sum += nodes[i].right.val
                nums++
            }
        }
        nodes = nodes.slice(length)
        if(nums !== 0) {
            result.push(sum / nums)
        }
    }
    return result
};