func isIsomorphic(s string, t string) bool {
    return mapping(s, t) && mapping(t, s) //important 正映射只能保证s中相同字母映射中t相同字母，"eb", "bb"
}

func mapping(s string, t string) bool {
    letterMapping := map[byte]byte{}
    for i:=0; i<len(s); i++ {
        value, has := letterMapping[s[i]]
        if has == true {
            if value != t[i] {
                return false
            }
        } else {
            letterMapping[s[i]] = t[i]
        }
    }
    return true
}