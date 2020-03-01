func findLHS(nums []int) int {
    mapping := map[int]int{}
    max := 0
    for i:=0; i<len(nums); i++ {
        count, has := mapping[nums[i]]
        if has {
            mapping[nums[i]] = count + 1
        } else {
            mapping[nums[i]] = 1
        }
    }
    for i:=0; i<len(nums); i++ {
        greaterCount, hasGreater := mapping[nums[i]+1]
        lessCount, hasLess := mapping[nums[i]-1]
        if hasGreater {
            if greaterCount + mapping[nums[i]] > max {
                max = greaterCount + mapping[nums[i]]
            }
        }
        if hasLess {
            if lessCount + mapping[nums[i]] > max {
                max = lessCount + mapping[nums[i]]
            }
        }
    }
    return max
}