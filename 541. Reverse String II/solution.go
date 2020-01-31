func reverseStr(s string, k int) string {
    arr := []rune(s)
    for i:=0; i<len(s); i=i+k*2 {
        start := i
        end := int(math.Min(float64(i+k-1), float64(len(s)-1)))
        for {
            if start >= end {
                break
            }
            temp := arr[end]
            arr[end] = arr[start]
            arr[start] = temp
            start++
            end--
        }
    }
    return string(arr)
}

func reverseAllStr(s string) string {
    arr := []rune(s)
    for i:=0; i<len(arr)/2; i++ {
        temp := arr[len(arr)-1-i]
        arr[len(arr)-1-i] = arr[i]
        arr[i] = temp
    }
    return string(arr)
}