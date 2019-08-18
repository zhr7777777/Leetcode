func addBinary(a string, b string) string {
    loopLength := 0
    rest := ""
    if len(a) > len(b) {
        loopLength = len(b)
        rest = a[0:len(a)-loopLength]
    } else {
        loopLength = len(a)
        rest = b[0:len(b)-loopLength]
    }
    result := ""
    flag := 0
    for i:=0; i<loopLength; i++ {
        bitA, errA := strconv.Atoi(string(a[len(a)-1-i]))
        bitB, errB := strconv.Atoi(string(b[len(b)-1-i]))
        if errA == nil && errB == nil {
            result = strconv.Itoa(flag ^ bitA ^ bitB) + result
            if bitA + bitB + flag >= 2 {
                flag = 1
            } else {
                flag = 0
            }
        }
    }
    if flag == 1 {
        rest = addBinary(rest, "1")
    }
    result = rest + result
    return result
}