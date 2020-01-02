func detectCapitalUse(word string) bool {
    if len(word) == 1 {
        return true
    }
    if word[0] >= 'A' && word[0] <= 'Z' {
        if word[1] >= 'A' && word[1] <= 'Z' {
            for i:=2; i<len(word); i++ {
                if word[i] < 'A' || word[i] > 'Z' {
                    return false
                }
            }
        } else {
            for i:=2; i<len(word); i++ {
                if word[i] >= 'A' && word[i] <= 'Z' {
                    return false
                }
            }
        }
    } else {
        for i:=1; i<len(word); i++ {
            if word[i] < 'a' || word[i] > 'z' {
                return false
            }
        }   
    }
    return true
}