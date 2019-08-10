func removeElement(nums []int, val int) int {
    var curIndex int = 0
    for i:=0; i<len(nums); i++ {
        if nums[i] != val {
            nums[curIndex] = nums[i]
            curIndex++
        }
    }
    return curIndex
}