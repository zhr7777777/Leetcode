/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function(s) {
    let max = 0
    let record = ''
    for(let i=0; i<s.length; i++) {
        if(!record.includes(s[i])) {
            record += s[i]
        } else {
            max = Math.max(max, record.length)
            let index = record.indexOf(s[i])
            record = index+1 < record.length ? record.slice(index+1) : ''
            record += s[i]  // !!!
        }
    }
    return Math.max(max, record.length) // !!!
};