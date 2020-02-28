/**
 * // Definition for a Node.
 * function Node(val, children) {
 *    this.val = val;
 *    this.children = children;
 * };
 */
/**
 * @param {Node} root
 * @return {number[]}
 */
var preorder = function(root) {
    var number = []
    var perTraverse = function (root) {
        if(!root) {
            return []
        }
        number.push(root.val)
        if(root.children) {
            root.children.forEach(e => {
                perTraverse(e)
            })
        }
        return number
    }
    return perTraverse(root)
};