func singleNumber(nums []int) int {
    var numsMap map[int]int
    result := 0
    numsMap = make(map[int]int)
    for i:=0; i<len(nums); i++ {
        index, has := numsMap[nums[i]]
        if has {
            delete(numsMap, nums[i])
        } else {
            numsMap[nums[i]] = index
        }
    }
    for key := range numsMap {
        result = key
    }
    return result
}