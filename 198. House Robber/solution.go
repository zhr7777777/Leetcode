func rob(nums []int) int {
    length := len(nums)
    if length == 0 {
        return 0
    }
    if length == 1 {
        return nums[0]
    }
    a := nums[0]
    b := nums[1]
    record := []int{a, max(a, b)}
    for i:=2; i<length; i++ {
        record = append(record, max(record[len(record)-1], record[len(record)-2] + nums[i]))
    }
    return max(record[len(record)-1], record[len(record)-2])
}

func max(a int, b int) int {
    if a >= b {
        return a
    }
    return b
}