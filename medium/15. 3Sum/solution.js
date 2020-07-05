/**
 * @param {number[]} nums
 * @return {number[][]}
 */
var threeSum = function(nums) {
    nums = nums.sort((a, b) => a - b)   // 升序排列
    let result = []
    for(let i=0; i<nums.length-2; i++) {
        if(i===0 || (i > 0 && nums[i] !== nums[i-1])) { // 当前元素和上一个元素相等，找过了跳过
            let start = i + 1
            let end = nums.length - 1
            let target = 0 - nums[i]
            while(end > start) {
                if(start > i+1 && nums[start] === nums[start-1]) { // 起始元素和上一个起始元素相同，两个元素确定了第三个元素，避免重复，应跳过
                    start = start + 1
                    continue
                }
                if(nums[start] + nums[end] === target) {
                    result.push([nums[i], nums[start], nums[end]])
                    start++ // 需要继续找
                } else if(nums[start] + nums[end] > target) {
                    end--
                } else {
                    start++
                }
            }
        }
    }
    return result
};