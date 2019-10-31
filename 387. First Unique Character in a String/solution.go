func firstUniqChar(s string) int {
    index := len(s)
    mapping := map[byte]int{}
    for i:=0; i<len(s); i++ {
        _, has := mapping[s[i]]
        if has {
            mapping[s[i]] = len(s)
        } else {
            mapping[s[i]] = i
        }
    }
    for c := range mapping {
        if mapping[c] < index {
            index = mapping[c]
        }
    }
    if index == len(s) {
        return -1
    }
    return index
}