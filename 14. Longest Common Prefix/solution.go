func longestCommonPrefix(strs []string) string {
    if(len(strs) == 0) {
        return ""
    }
    if(len(strs) == 1) {
        return strs[0]
    }
    var cur string = strs[0]
    var result string = ""
    for i := 1; i < len(strs); i++ {
        result = ""
        for j := 0; j < min(len(cur), len(strs[i])); j++ {
            if string(cur[j]) == string(strs[i][j]) {
                result += string(strs[i][j])
            } else {
                break
            }
        }
        if result == "" {
            break
        }
        cur = result
    }
    return result
}

func min(a int, b int) int {
    if a >= b {
        return b
    }
    return a
}