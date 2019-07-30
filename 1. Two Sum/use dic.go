func twoSum(nums []int, target int) []int {
    var numsMap map[int]int
    var result [] int
    numsMap = make(map[int]int)
    numsMap[nums[0]] = 0
    for i := 1; i<len(nums); i++ {
        var search int = target - nums[i]
        idx, has := numsMap[search]
        if has {
            result = append(result, idx)
            result = append(result, i)
            break
        }
        numsMap[nums[i]] = i
    }
    return result
}