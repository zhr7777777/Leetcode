/**
 * @param {number[]} nums
 * @return {number}
 */
// dp[0] = nums[0]
// dp[1] = max(nums[0], nums[1])
// dp[i] = max(dp[i-1], dp[i-2] + nums[i])
var rob = function(nums) {
    if(nums.length === 0) return 0
    if(nums.length === 1) return nums[0]
    if(nums.length === 2) return Math.max(nums[0], nums[1])
    let records = [nums[0],  Math.max(nums[0], nums[1])]
    for(let i=2; i<nums.length; i++) {
        records.push(Math.max(records[i-1], records[i-2]+nums[i]))
    }
    return records[records.length-1]
};