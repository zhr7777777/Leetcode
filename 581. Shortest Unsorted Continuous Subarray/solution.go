func findUnsortedSubarray(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    l := 0
    r := len(nums) - 1
    shouldSort := false
    min := 1000
    max := -1000
    for i:=0; i<len(nums)-1; i++ {
        if nums[i+1] < nums[i] {
            shouldSort = true
            if nums[i+1] < min {
                min = nums[i+1]
            }
        }
    }
    if !shouldSort {
        return 0
    } else {
        for i:=len(nums)-1; i>0; i-- {
            if nums[i-1] > nums[i] {
                if nums[i-1] > max {
                    max = nums[i-1]
                }
            }
        }
        for ;l<len(nums); l++ {
            if min < nums[l] {
                break
            }
        }
        for ;r>-1; r-- {
            if max > nums[r] {
                break
            }
        }
    }
    return r - l + 1
}

func min(a int, b int) int {
    return int(math.Min(float64(a), float64(b)))
}