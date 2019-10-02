func containsNearbyDuplicate(nums []int, k int) bool {
    mapping := map [int]int{}
    for i:=0; i<len(nums); i++ {
        index, has := mapping[nums[i]]
        if has {
            if i - index <= k {
                return true
            } else {
                mapping[nums[i]] = i
            }
        } else {
            mapping[nums[i]] = i
        }
    }
    return false
}
