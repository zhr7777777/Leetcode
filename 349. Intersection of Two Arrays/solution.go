func intersection(nums1 []int, nums2 []int) []int {
    mapping := map[int]int{}
    result := []int{}
    for i:=0; i<len(nums1); i++ {
        _, has := mapping[nums1[i]]
        if !has {
            mapping[nums1[i]] = 0
        }
    }
    for i:=0; i<len(nums2); i++ {
        sign, has := mapping[nums2[i]]
        if has && sign == 0 {
            mapping[nums2[i]] = -1
            result = append(result, nums2[i])
        }
    }
    return result
}