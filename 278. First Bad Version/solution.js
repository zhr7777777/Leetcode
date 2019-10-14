/**
 * Definition for isBadVersion()
 * 
 * @param {integer} version number
 * @return {boolean} whether the version is bad
 * isBadVersion = function(version) {
 *     ...
 * };
 */

/**
 * @param {function} isBadVersion()
 * @return {function}
 */
var solution = function(isBadVersion) {
    /**
     * @param {integer} n Total versions
     * @return {integer} The first bad version
     */
    return function(n) {
        if(isBadVersion(1)) {
            return 1
        }
        let first = 1
        let end = n
        while(end - first > 1) {
            let middle = parseInt((end + first) / 2)
            if(isBadVersion(middle)) {
                end = middle
            } else {
                first = middle
            }
        }
        return end
    };
};