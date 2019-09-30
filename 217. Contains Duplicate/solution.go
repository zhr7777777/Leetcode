func containsDuplicate(nums []int) bool {
    mapping := map [int]int{}
    for i:=0; i<len(nums); i++ {
        _, has := mapping[nums[i]]
        if has {
            return true
        } else {
            mapping[nums[i]] = nums[i]
        }
        
    }
    return false
}