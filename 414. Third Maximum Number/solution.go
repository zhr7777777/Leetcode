func thirdMax(nums []int) int {
    diffThree := []int{}
    for i:=0; i<len(nums); i++ {
        if len(diffThree) == 3 {
            break
        }
        for j:=0; j<len(nums)-i-1; j++ {
            if nums[j] > nums[j+1] {
                temp := nums[j+1]
                nums[j+1] = nums[j]
                nums[j] = temp
            }
        }
        has := false
        for j:=0; j<len(diffThree); j++ {
            if nums[len(nums)-1-i] == diffThree[j] {
                has = true
                break
            }
        }
        if !has {
            diffThree = append(diffThree, nums[len(nums)-1-i])
        }
    }
    if len(diffThree) == 1 {
        return diffThree[0]
    }
    if len(diffThree) == 2 {
        return diffThree[0]
    }
    return diffThree[2]
}

// func thirdMax(nums []int) int {
//     if len(nums) == 1 {
//         return nums[0]
//     }
//     diffThree := []int{}
//     diffThree = append(diffThree, nums[0])
//     i := 1
//     for ; i<len(nums); i++ {
//         if len(diffThree) == 3 {
//             break
//         }
//         has := false
//         for j:=0; j<len(diffThree); j++ {
//             if nums[i] == diffThree[j] {
//                 has = true
//                 break
//             }
//         }
//         if !has {
//             diffThree = append(diffThree, nums[i])
//         }
//     }
//     sortedThree := sort(diffThree)
//     if len(diffThree) == 1 {
//         return sortedThree[0]
//     }
//     if len(diffThree) == 2 {
//         return sortedThree[1]
//     }
//     thirdMax := sortedThree[0]
//     secondMax := sortedThree[1]
//     max := sortedThree[2]
//     for ; i<len(nums); i++ {
//         if nums[i] == max || nums[i] == secondMax || nums[i] == thirdMax {  // important
//             continue
//         }
//         if nums[i] > max {
//             thirdMax = secondMax
//             secondMax = max
//             max = nums[i]
//             continue
//         }
//         if nums[i] > secondMax {
//             thirdMax = secondMax
//             secondMax = nums[i]
//             continue
//         }
//         if nums[i] > thirdMax {
//             thirdMax = nums[i]
//         }
//     }
//     return thirdMax
// }

// func sort(nums []int) []int {
//     for i:=0; i<len(nums); i++ {
//         for j:=0; j<len(nums)-i-1; j++ {
//             if nums[j] > nums[j+1] {
//                 temp := nums[j+1]
//                 nums[j+1] = nums[j]
//                 nums[j] = temp
//             }
//         }
//     }
//     return nums
// }