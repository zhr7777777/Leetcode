func licenseKeyFormatting(S string, K int) string {
    k := K
    s := ""
    for i:=0; i<len(S); i++ {
        cur := rune(S[len(S)-1-i])
        if k == 0 {
            k = K
            s = "-" + s
        }
        if cur == '-' {
            
        } else {
            if k == 0 {
                k = K
                s = "-" + s
            } 
            if unicode.IsLower(cur) { // important
                cur = unicode.ToUpper(cur)
            }
            s = string(cur) + s
            k--
        }
    }
    if len(s) > 0 && s[0] == '-' {  // important
        return s[1:]
    }
    return s
}