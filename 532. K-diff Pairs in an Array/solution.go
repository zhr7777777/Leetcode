func findPairs(nums []int, k int) int {
    if k < 0 {  // important
        return 0
    }
    pairMap := map[int]int{}
    numsMap := map[int]int{}
    count := 0
    for i:=0; i<len(nums); i++ {
        _, has := numsMap[nums[i]]
        if !has {
            numsMap[nums[i]] = 0
            c, hasPair := pairMap[nums[i]]
            if hasPair {
                count = count + c
            }
            leftCount, hasLeft := pairMap[nums[i]-k]
            if hasLeft {    // important
                pairMap[nums[i]-k] = leftCount + 1
            } else {
                pairMap[nums[i]-k] = 1
            }
            rightCount, hasRight := pairMap[nums[i]+k]
            if hasRight {   // important
                pairMap[nums[i]+k] = rightCount + 1
            } else {
                pairMap[nums[i]+k] = 1
            }
        } else {
            if k == 0 && numsMap[nums[i]] == 0 { // important
                count++
                numsMap[nums[i]] = -1
            }
        }
    }
    return count
}