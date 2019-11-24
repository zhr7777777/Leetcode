func compress(chars []byte) int {
    if len(chars) == 0 {
        return 0
    }
    cur := chars[0]
    str := ""
    count := 1
    for i:=1; i<len(chars); i++ {
        if chars[i] == cur {
            count++
        } else {
            str = str + string(cur)
            if count != 1 {
                str = str + strconv.Itoa(count)
            }
            count = 1
            cur = chars[i]
        }
    }
    str = str + string(cur) // important
    if count != 1 {
        str = str + strconv.Itoa(count)
    }
    for i:=0; i<len(str); i++ {
        chars[i] = str[i]
    }
    chars = chars[:len(str)]
    return len(str)
}