func findTheDifference(s string, t string) byte {
    mapping := map[byte]int{}
    for i:=0; i<len(s); i++ {
        count, has := mapping[s[i]]
        if !has {
            mapping[s[i]] = 1
        } else {
            mapping[s[i]] = count + 1   
        }
    }
    for i:=0; i<len(t); i++ {
        _, has := mapping[t[i]]
        if !has {
            return t[i]
        }
        mapping[t[i]] = mapping[t[i]] - 1
        if mapping[t[i]] < 0 {
            return t[i]
        }
    }
    return s[0]
}