func rotate(nums []int, k int) {   // 236ms
    if len(nums) == 0 {
        return
    }
    k = k % len(nums)
    if k == 0 {
        return
    }
    for i:=0; i<k; i++ {
        last := nums[len(nums)-1]
        for j:=0; j<len(nums); j++ {
            if j == len(nums) - 1 {
                nums[0] = last
                break
            }
            nums[len(nums)-1-j] = nums[len(nums)-1-j-1]
        }
    }
}