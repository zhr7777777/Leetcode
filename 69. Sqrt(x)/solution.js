/**
 * @param {number} x
 * @return {number}
 */
var mySqrt = function(x) {  // 使用二分查找
    if(x < 2) return x
    let left = 1
    let right = parseInt(x / 2)
    while(left <= right) {
        let mid = left + parseInt((right - left) / 2)
        let mul = mid * mid
        if(mul === x) return mid
        if(mul < x) {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return right
};