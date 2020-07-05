/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @param {TreeNode} p
 * @param {TreeNode} q
 * @return {TreeNode}
 */
var lowestCommonAncestor = function(root, p, q) {
    let [p1, p2] = findPaths(root, p, q)
    for(let i=0; i<p1.length; i++) {
        let cur = p1[p1.length-i-1]
        for(let j=0; j<p2.length; j++) {
            if(cur === p2[p2.length-j-1]) {
                return cur
            }
        }
    }
    return null
};

var findPaths = function(root, p, q) {
    let stack = []
    let cur = root
    let paths = []
    while(stack.length || cur) {
        if(cur) {
            stack.push(cur)
            cur = cur.left
        } else {
            let top = stack[stack.length-1]
            if(!top.visited) {
                top.visited = true
                cur = top.right
            } else {
                if(top === p || top === q) {
                    console.log(stack.map(e => e.val).join('->'))
                    paths.push([...stack])
                }
                stack.pop()
            }
        }
    }
    return paths
}