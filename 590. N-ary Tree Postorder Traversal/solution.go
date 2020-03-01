/**
 * // Definition for a Node.
 * function Node(val,children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */
/**
 * @param {Node} root
 * @return {number[]}
 */
var postorder = function(root) {
    var numbers = []
    var postTraverse = function (root) {
        if(!root) {
            return []
        }
        if(root.children) {
            root.children.forEach(e => {
                postTraverse(e)
            })
        }
        numbers.push(root.val)
        return numbers
    }
    return postTraverse(root)
};