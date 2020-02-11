func arrayPairSum(nums []int) int {
    nums = bubbleSort(nums)
    sum := 0
    for i:=0; i<len(nums); i=i+2 {
        sum += nums[i]
    }
    return sum
}

func bubbleSort(nums []int) []int {
    for i:=0; i<len(nums)-1; i++ {
        for j:=0; j<len(nums)-i-1; j++ {
            if nums[j] > nums[j+1] {
                temp := nums[j+1]
                nums[j+1] = nums[j]
                nums[j] = temp
            }
        }
    }
    return nums
}