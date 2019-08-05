func isValid(s string) bool {
    if len(s) == 0 {
        return true
    }
    if len(s) % 2 == 1 {
        return false
    }
    stack := []string{string(s[0])}
    parenthesesMap := map[string]string{"(": ")", "{": "}", "[": "]"}
    for i:=1; i<len(s); i++ {
        if len(stack) == 0 {
            stack = append(stack, string(s[i]))
            continue
        }
        if parenthesesMap[stack[len(stack)-1]] == string(s[i]) {
            stack = stack[:len(stack)-1]
        } else {
            stack = append(stack, string(s[i]))
        }
    }
    if len(stack) == 0 {
        return true
    } else {
        return false
    }
}