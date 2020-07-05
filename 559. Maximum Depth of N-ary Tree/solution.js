/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */

/**
 * @param {Node} root
 * @return {number}
 */
var maxDepth = function(root) {
    if(!root) return 0
    let max = 0
    let queue = [root]
    while(queue.length) {
        let length = queue.length
        for(let i=0; i<length; i++) {
            let cur = queue.shift()
            if(cur.children.length) {
                queue = queue.concat(cur.children)
            }
        }
        max++
    }
    return max
};