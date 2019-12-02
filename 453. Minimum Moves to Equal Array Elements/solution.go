func minMoves(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    max := nums[0]
    maxIndex := 0
    min := nums[0]
    sum := 0
    moves := 0
    for i:=0; i<len(nums); i++ {
        if nums[i] > max {
            maxIndex = i
            max = nums[i]
        }
        if nums[i] < min {
            min = nums[i]
        }
        sum += nums[i]
    }
    for {
        if sum % len(nums) == 0 && sum / len(nums) == max {
            return moves
        }
        curMoves := max - min
        curMaxIndex := maxIndex
        min = max
        for i:=0; i<len(nums); i++ {
            if i == maxIndex {
                continue
            }
            nums[i] = nums[i] + curMoves
            if nums[i] > max {
                curMaxIndex = i
                max = nums[i]
            }
            sum += curMoves
        }
        maxIndex = curMaxIndex
        moves += curMoves
    }
}