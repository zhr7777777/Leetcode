func lengthOfLastWord(s string) int {
    result := 0
    isLast := true
    for i:=0; i<len(s); i++ {
      if string(s[len(s)-1-i]) == " " {
          if isLast == true {
              continue
          } else {
              break
          }
      } else {
          isLast = false
          result++
      }
    }
    return result
}