/**
 * @param {string} s
 * @return {string}
 */
var reverseWords = function(s) {
    let matches = s.match(/[^\s]+/g)
    return Array.isArray(matches) ? matches.reverse().join(' ') : ''
};