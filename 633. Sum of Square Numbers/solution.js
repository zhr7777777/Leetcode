/**
 * @param {number} c
 * @return {boolean}
 */
var judgeSquareSum = function(c) {
    if(c === 0) return true
    let threshold = Math.sqrt(c)
    for(let a=0; a<threshold; a++) {
        if(isInteger(Math.sqrt(c - a * a))) {
            return true
        }
    }
    return false
};

var isInteger = num => {
    return parseInt(num) === num
}