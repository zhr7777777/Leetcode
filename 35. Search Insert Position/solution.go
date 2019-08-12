func searchInsert(nums []int, target int) int {
    result := -1
    for i:=0; i<len(nums); i++ {
        if nums[i] >= target {
            result = i
            break
        }
    }
    if result == -1 {
        return len(nums)
    }
    return result
}