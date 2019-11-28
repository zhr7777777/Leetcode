func findDisappearedNumbers(nums []int) []int { // 404ms
    mapping := map[int]bool{}
    storeMapping := map[int]bool{}
    disappearNums := []int{}
    for i:=1; i<=len(nums); i++ {
        _, inStore := storeMapping[i]
        if !inStore {
            mapping[i] = true
        }
        _, hasFront := mapping[nums[i-1]]
        if hasFront {
            delete(mapping, nums[i-1])
        } else {
            storeMapping[nums[i-1]] = true
        }
    }
    for i := range mapping {
        disappearNums = append(disappearNums, i)
    }
    return disappearNums
}

func findDisappearedNumbers(nums []int) []int { // 396ms
    disappearNums := []int{}
    for i:=0; i<len(nums); i++ {
        for {
            if nums[i] == nums[nums[i]-1] { // important
                break
            }
            temp := nums[nums[i]-1]
            nums[nums[i]-1] = nums[i]
            nums[i] = temp
        }
    }
    for i:=0; i<len(nums); i++ {
        if nums[i] != i + 1 {
            disappearNums = append(disappearNums, i+1)
        }
    }
    return disappearNums
}