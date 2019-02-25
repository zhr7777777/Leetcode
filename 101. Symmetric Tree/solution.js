/**
 * Definition for a binary tree node.
 * function TreeNode(val) {
 *     this.val = val;
 *     this.left = this.right = null;
 * }
 */
/**
 * @param {TreeNode} root
 * @return {boolean}
 */
var isSymmetric = function(root) {
    // mid order traverse cannot be used
    return isMirror(root, root)
};
    
var isMirror = function(p, q) {
    if(!p && !q) return true
    if(!p && q) return false
    if(p && !q) return false
    if(p.val !== q.val) return false
    return isMirror(p.left, q.right) && isMirror(p.right, q.left) // compare left part with right part
}