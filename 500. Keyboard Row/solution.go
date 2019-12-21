func findWords(words []string) []string {
    keyRows := map[string]int{
        "q": 1, "w": 1, "e": 1, "r": 1, "t": 1, "y": 1, "u": 1, "i": 1, "o": 1, "p": 1,
        "a": 2, "s": 2, "d": 2, "f": 2, "g": 2, "h": 2, "j": 2, "k": 2, "l": 2,
        "z": 3, "x": 3, "c": 3, "v": 3, "b": 3, "n": 3, "m": 3}
    foundWords := []string{}
    for i:=0; i<len(words); i++ {
        row := keyRows[string(unicode.ToLower(rune(words[i][0])))]
        flag := true
        for j:=1; j<len(words[i]); j++ {
            if keyRows[string(unicode.ToLower(rune(words[i][j])))] != row {
                flag = false
                break
            }
        }
        if flag {
            foundWords = append(foundWords, words[i])
        }
    }
    return foundWords
}