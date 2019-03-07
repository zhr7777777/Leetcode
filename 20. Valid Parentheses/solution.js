/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function(s) {
    if(s.length % 2 === 1) {
        return false
    }
    let stack = []
    let map = {
        '{': '}',
        '(': ')',
        '[': ']'
    }
    for(let c of s) {
        if(!stack.length) {
            stack.push(c)
        } else {
            let top = stack[stack.length - 1]
            if(map[top] === c) {
                stack.pop()
            } else {
                stack.push(c)
            }
        }
    }
    return !Boolean(stack.length) 
};