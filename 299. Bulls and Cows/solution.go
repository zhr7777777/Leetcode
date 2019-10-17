func getHint(secret string, guess string) string {
    bulls := 0
    cows := 0
    mapping := map[byte]int{}
    newGuess := ""
    for i:=0; i<len(secret); i++ {
        if secret[i] == guess[i] {
            bulls++
        } else {
            newGuess = newGuess + string(guess[i])
            count, has := mapping[secret[i]]
            if has {
                mapping[secret[i]] = count + 1  // important
            } else {
                mapping[secret[i]] = 1
            }
        }
    }
    for i:=0; i<len(newGuess); i++ {
        count, has := mapping[newGuess[i]]
        if has {
            if count > 0 {  // important
                cows++
                mapping[newGuess[i]] = count - 1   
            }
        }
    }
    return strconv.Itoa(bulls) + "A" + strconv.Itoa(cows) + "B"
}