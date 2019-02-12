/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {string[]}
 */
var binaryTreePaths = function(root) {
    var cur = root
    var stack = []
    var result = []
    while(cur || stack.length) {
        if(cur) {
            cur.visited = false
            stack.push(cur)
            cur = cur.left
        } else {
            let top = stack[stack.length - 1]
            if(top.visited) {
                if(!top.right && !top.left) {
                    result.push(stack.map(e => e.val).join('->'))
                }
                let output = stack.pop()
            } else {
                top.visited = true
                cur = top.right
            }
        }
    }
    return result
};