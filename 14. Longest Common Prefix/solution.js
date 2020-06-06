/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function(strs) {
    if(!strs.length) return ""
    if(strs.length === 1) return strs[0]
    let curCommon = strs[0]
    for(let i=1; i<strs.length; i++) {
        let common = ''
        let length = Math.min(curCommon.length, strs[i].length)
        for(let j=0; j<length; j++) {
            if(curCommon[j] === strs[i][j]) {
                common = common + strs[i][j]
            } else {
                break   // !!!
            }
        }
        curCommon = common
        if(!curCommon) return ''
    }
    return curCommon
};