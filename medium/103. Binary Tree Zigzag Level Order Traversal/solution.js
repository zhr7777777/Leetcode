/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number[][]}
 */
var zigzagLevelOrder = function(root) {
    if(!root) return []
    const result = []
    let queue = [root]
    let level = 0
    while(queue.length) {
        const levelNodes = []
        const length = queue.length
        for(let i=0; i<length; i++) {
            levelNodes.push(queue[i].val)
            if(queue[i].left) queue.push(queue[i].left)  
            if(queue[i].right) queue.push(queue[i].right)
        }
        result.push(level % 2 === 0 ? levelNodes : levelNodes.reverse())
        level++
        queue = queue.slice(length)
    }
    return result
};