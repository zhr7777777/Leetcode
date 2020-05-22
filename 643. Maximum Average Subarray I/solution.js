/**
 * @param {number[]} nums
 * @param {number} k
 * @return {number}
 */
var findMaxAverage = function(nums, k) {
    if(nums.length === k) return sumArr(nums) / k
    let times = nums.length - k
    let max = sumArr(nums.slice(0, k))
    for(let i=0; i<times; i++) {
        let cur = sumArr(nums.slice(i+1, i+1+k))
        if(cur > max) {
            max = cur
        }
    }
    return max / k
};

var sumArr = function(nums) {
    return nums.reduce((sum, num) => sum + num, 0)
}