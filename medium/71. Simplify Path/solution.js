/**
 * @param {string} path
 * @return {string}
 */
var simplifyPath = function(path) {
    let paths = path.split('/').filter(p => p !== '')
    let stack = []
    for(let i=0; i<paths.length; i++) {
        if(paths[i] === '..') {
            stack.pop()
        } else if (paths[i] === '.') {
            continue    
        } else {
            stack.push(paths[i])
        }
    }
    return '/' + stack.join('/')
};