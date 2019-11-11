func longestPalindrome(s string) int {
    mapping := map[byte]int{}
    length := 0
    extra := 0
    for i:=0; i<len(s); i++ {
        count, has := mapping[s[i]]
        if has {
            mapping[s[i]] = count + 1
        } else {
            mapping[s[i]] = 1
        }
    }
    for c := range mapping {
        count := mapping[c] 
        if count % 2 == 0 {
            length += count
        } else {
            length = length + count - 1 // important
            extra = 1
        }
    }
    return length + extra
}