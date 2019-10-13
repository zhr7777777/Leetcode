func missingNumber(nums []int) int {
    arr := make([]int, len(nums)+1)
    missingNumber := -1
    for i:=0; i<len(arr); i++ {
        arr[i] = -1
    }
    for i:=0; i<len(nums); i++ {
        arr[nums[i]] = nums[i]
    }
    for i:=0; i<len(arr); i++ {
        if arr[i] == -1 {
            missingNumber = i
            break
        }
    }
    return missingNumber
}