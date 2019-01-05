/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(nums, target) {
    var result = [];
    outside:
    for(var i=0; i<nums.length; i++) {
        let search = target - nums[i];
        for(var j=i+1; j< nums.length; j++) {
            if(nums[j] == search) {
                result[0] = i;
                result[1] = j;
                break outside;
            }
        }
    }
    return result;
};