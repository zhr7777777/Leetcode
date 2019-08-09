func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    var cur int = nums[0]
    var curIndex = 1
    for i:=0; i<len(nums); i++ {
        if i == 0 {
            continue
        }
        if nums[i] != cur {
            nums[curIndex] = nums[i]
            curIndex++
            cur = nums[i]
        }
    }
    return len(nums[:curIndex])
}