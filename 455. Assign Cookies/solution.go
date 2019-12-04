func findContentChildren(g []int, s []int) int {
    if len(g) == 0 || len(s) == 0 {
        return 0
    }
    gSorted := sort(g)
    sSorted := sort(s)
    curgIndex := 0
    cursIndex := 0
    max := 0
    for {
        if curgIndex == len(g) || cursIndex == len(s) {
            break
        }
        if sSorted[len(sSorted)-1] < gSorted[curgIndex] {
            break
        }
        if sSorted[cursIndex] >= gSorted[curgIndex] {
            curgIndex++
            max++
        }
        cursIndex++
    }
    return max
}

func sort(nums []int) []int {
    for i:=0; i<len(nums)-1; i++ {
        for j:=0; j<len(nums)-i-1; j++ {
            if nums[j] > nums[j+1] {
                temp := nums[j]
                nums[j] = nums[j+1]
                nums[j+1] = temp
            }
        }
    }
    return nums
}