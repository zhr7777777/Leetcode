func maximumProduct(nums []int) int {
    length := len(nums)
    for i:=0; i<3; i++ {
        for j:=0; j<length-i-1; j++ {
            if nums[j] > nums[j+1] {
                temp := nums[j]
                nums[j] = nums[j+1]
                nums[j+1] = temp
            }
        }
    }
    if length > 3 {
        for i:=0; i<2; i++ {
            for j:=length-4; j>i; j-- {
                if nums[j] < nums[j-1] {
                    temp := nums[j]
                    nums[j] = nums[j-1]
                    nums[j-1] = temp
                }
            }
        }   
    }
    a := nums[length-1] * nums[length-2] * nums[length-3]
    b := nums[0] * nums[1] * nums[length-1]
    if a < b {
        return b
    }
    return a
}