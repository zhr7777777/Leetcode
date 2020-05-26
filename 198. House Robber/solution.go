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
    return record[len(record)-1]
}

func max(a int, b int) int {
    if a >= b {
        return a
    }
    return b
}

// 一开始思路是从0开始，或者从1开始，比较每次跳一个或两个，取最大，最后两个出发点取最大
// 反例：[2,4,8,9,9,3]
// 应该使用贪心算法