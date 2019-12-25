func findRelativeRanks(nums []int) []string {
    rank := make([]string, len(nums))
    for i:=0; i<len(nums); i++ {
        max := i
        for j:=0; j<len(nums); j++ {
            if nums[j] > nums[max] {
                max = j
            }
        }
        switch i {
            case 0:
                rank[max] = "Gold Medal"
                break
            case 1:
                rank[max] = "Silver Medal"
                break
            case 2:
                rank[max] = "Bronze Medal"
                break
            default:
                rank[max] = strconv.Itoa(i+1)
                break
        }
        nums[max] = -1
    }
    return rank
}