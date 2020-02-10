func reverseWords(s string) string {
    words := strings.Split(s, " ")
    for i:=0; i<len(words); i++ {
        words[i] = reverseWord(words[i])
    }
    return strings.Join(words, " ")
}

func reverseWord(s string) string {
    arr := []rune(s)
    for i:=0; i<len(s)/2; i++ {
        temp := arr[len(arr)-1-i]
        arr[len(arr)-1-i] = arr[i]
        arr[i] = temp
    }
    return string(arr)
}