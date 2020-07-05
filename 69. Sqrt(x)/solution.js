/**
 * @param {number} x
 * @return {number}
 */
// var mySqrt = function(x) {  // 使用二分查找
//     if(x < 2) return x
//     let left = 1
//     let right = parseInt(x / 2)
//     while(left <= right) {
//         let mid = left + parseInt((right - left) / 2)
//         let mul = mid * mid
//         if(mul === x) return mid
//         if(mul < x) {
//             left = mid + 1
//         } else {
//             right = mid - 1
//         }
//     }
//     return right
// };

const mySqrt = num => {
    if(num < 2) return num
    let start = 1
    let end = parseInt(num / 2)
    while(end >= start) {
        let mid = start + parseInt((end - start) / 2)
        let mul = mid * mid
        if(mul === num) return mid
        if(mul > num) {
            end = mid - 1
        } else {
            start = mid + 1
        }
    }
    return end // 需要返回end，因为start = mid + 1后可能循环终止，最终start * start > num
}

console.log(mySqrt(10))