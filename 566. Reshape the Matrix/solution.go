func matrixReshape(nums [][]int, r int, c int) [][]int {
    if len(nums) == 0 {
        return nums
    }
    if len(nums) * len(nums[0]) != r * c {
        return nums
    }
    if len(nums) == r && len(nums[0]) == c {
        return nums
    }
    tempM := []int{}
    newM := [][]int{}
    for i:=0; i<len(nums); i++ {
        for j:=0; j<len(nums[i]); j++ {
            tempM = append(tempM, nums[i][j])   
        }
    }
    for i:=0; i<r; i++ {
        newM = append(newM, tempM[i*c:(i+1)*c])
    }
    return newM
}