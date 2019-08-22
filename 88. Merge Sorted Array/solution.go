func merge(nums1 []int, m int, nums2 []int, n int)  {
    if len(nums2) == 0 {
        return
    }
    j := 0
    for {
        if j == m {
            for i:=0; i<n; i++ {
                nums1[m+i] = nums2[i]
            }
            break
        }
        if nums2[0] < nums1[j] {
            temp := nums1[j]
            nums1[j] = nums2[0]
            nums2[0] = temp
            nums2 = resort(nums2)
        }
        j++
    }
}

func resort(arr []int) []int {
    for i:=0; i<len(arr); i++ {
        if i + 1 == len(arr) {
            break
        }
        if arr[i] > arr[i+1] {
            temp := arr[i+1]
            arr[i+1] = arr[i]
            arr[i] = temp
        } else {
            break
        }
    }
    return arr
} 
