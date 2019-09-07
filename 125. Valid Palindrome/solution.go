func isPalindrome(s string) bool {
    result := true
    leftIndex := 0
    rightIndex := len(s) - 1
    length := len(s)
    for i:=0; i<length/2; i++ {
        for {
            if leftIndex+i == len(s) { // important
                break
            }
            if isAlphanumeric(s[leftIndex+i]) {
                break
            }
            leftIndex++
            length--
        }
        for {
            if rightIndex-i == 0 {  // important
                break
            }
            if isAlphanumeric(s[rightIndex-i]) {
                break
            }
            rightIndex--
            length--
        }
        if leftIndex >= rightIndex {    // important
            break
        }
        if toUpper(s[leftIndex+i]) != toUpper(s[rightIndex-i]) {
            result = false
            break
        }
    }
    return result
}
                           
func isAlphanumeric(s byte) bool {
    result := false
    if 64 < s && s < 91 {
        result = true
    }
    if 96 < s && s < 123 {
        result = true
    }
    if 48 <= s && s <= 57 {
        result = true
    }
    return result
}
                           
func toUpper(s byte) byte {
    if 64 < s && s < 91 {
        s += 32
    }
    return s
}