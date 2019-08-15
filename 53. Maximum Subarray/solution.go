func maxSubArray(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    maxSum := nums[0]
    sum := nums[0]
    for i:=1; i<len(nums); i++ {
        if maxSum + nums[i] > nums[i] {
            maxSum = maxSum + nums[i]
        } else {
            maxSum = nums[i]
        }
        if maxSum >= sum {
            sum = maxSum
        }
    }
    return sum
}