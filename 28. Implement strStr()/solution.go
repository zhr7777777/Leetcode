func strStr(haystack string, needle string) int {
    if haystack == needle || needle == "" {
        return 0
    }
    if len(haystack) < len(needle) {
        return -1
    }
    var result int = -1
    for i:=0; i<len(haystack); i++ {
        if string(haystack[i]) == string(needle[0]) {
            if i+len(needle) <= len(haystack) && haystack[i:i+len(needle)] == needle {
                result = i
                break
            }
        }
    
    }
    return result
}