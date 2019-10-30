func canConstruct(ransomNote string, magazine string) bool {
    if len(magazine) < len(ransomNote) {
        return false
    }
    if strings.Contains(magazine, ransomNote) {
        return true
    }
    letterMapping := map[byte]int{}
    for i:=0; i<len(magazine); i++ {
        count, has := letterMapping[magazine[i]]
        if has {
            letterMapping[magazine[i]] = count + 1
        } else {
            letterMapping[magazine[i]] = 1
        }
    }
    for i:=0; i<len(ransomNote); i++ {
        count, has := letterMapping[ransomNote[i]]
        if has {
            if count == 0 {
                return false
            }
            letterMapping[ransomNote[i]] = count - 1
        } else {
            return false
        }
    }
    return true
}