/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
    if(nums.length === 0) return 0
    if(nums.length === 1) return nums[0]
    if(nums.length === 2) return Math.max(nums[0], nums[1])
    
    let arr = [nums[0], Math.max(nums[0], nums[1])] //局部最优 -> 全局最优，arr存储偷取每一家的最优解
    for(let i=2; i<nums.length; i++) {
        arr.push(Math.max(nums[i] + arr[i-2], arr[i-1])) // 基于上一步的最优解，比较偷当前house，不偷当前house，选择最大值
    }
    return arr[arr.length-1]
};