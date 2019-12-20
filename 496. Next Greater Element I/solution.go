func nextGreaterElement(nums1 []int, nums2 []int) []int {
    result := []int{}
    mapping := map[int]int{}
    for i:=0; i<len(nums2); i++ {
        mapping[nums2[i]] = i
    }
    for i:=0; i<len(nums1); i++ {
        index, _ := mapping[nums1[i]]
        flag := false
        j:= index+1
        for ; j<len(nums2); j++ {
            if nums2[j] > nums1[i] {
                flag = true
                break
            }
        }
        if flag {
            result = append(result, nums2[j])
        } else {
            result = append(result, -1)
        }
    }
    return result
}