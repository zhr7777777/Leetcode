/**
 * @param {number[]} nums
 * @return {number[]}
 */
var findErrorNums = function(nums) {
    let records = new Array(nums.length).fill(1)
    let result = []
    for(let i=0; i<nums.length; i++) {
        if(records[nums[i]-1] === 0) {
            result.push(nums[i])
        } else {
            records[nums[i]-1] = 0
        }
    }
    result.push(records.findIndex(num => num === 1) + 1)
    return result
};