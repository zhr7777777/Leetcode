func repeatedSubstringPattern(s string) bool {
    // arr := strings.Split(s, string(s[0]))
    // if len(arr) <= 2 {
    //     return false
    // }
    // cur := arr[1]
    // for i:=2; i<len(arr); i++ {
    //     if arr[i] != cur {
    //         return false
    //     }
    // }
    // return true
    l := len(s)
    for i:=1; i<=l/2; i++ {
        if checkStr(s, i) == true {
            return true
        }
    }
    return false
}

func checkStr(s string, n int) bool {
    if len(s) == 1 {
        return false
    }
    cur := s[:n]
    for i:=n; i<len(s); i=i+n {
        if i+n > len(s) {
            return false
        }
        if s[i:i+n] != cur {
            return false
        }
    }
    return true
}