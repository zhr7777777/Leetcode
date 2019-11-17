func countSegments(s string) int {
    arr := strings.Split(s, " ")
    result := len(arr)
    for i:=0; i<len(arr); i++ {
        if arr[i] == "" {
            result--
        }
    }
    return result
}