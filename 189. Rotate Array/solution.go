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

func rotate(nums []int, k int) {    // 52ms important
    if len(nums) == 0 {
        return
    }
    k = k % len(nums)
    if k == 0 {
        return
    }
    storage := append([]int(nil), nums[len(nums)-k:len(nums)]...)
    for i:=0; i<len(nums); i++ {
        if len(nums)-1-i-k < 0 {
            nums[len(nums)-1-i] = storage[len(nums)-1-i]
            continue
        }
        nums[len(nums)-1-i] = nums[len(nums)-1-i-k]
    }
}