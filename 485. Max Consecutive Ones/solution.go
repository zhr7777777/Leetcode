func findMaxConsecutiveOnes(nums []int) int {
    max := 0
    curMax := 0
    for i:=0; i<len(nums); i++ {
        if nums[i] == 0 {
            if curMax > max {
                max = curMax
            }
            curMax = 0
        } else {
            curMax++
            if curMax > max {   // important
                max = curMax
            }
        }
        if len(nums) - i - 1 < max && nums[i] == 0 { // [1,1,1,1,1,1,1,1]
            break
        }
    }
    return max
}