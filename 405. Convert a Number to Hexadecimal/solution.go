func toHex(num int) string {
    if num == 0 {
        return "0"
    }
    binaryStr := normalizeNum(num)
    hexStr := ""
    for i:=0; i<len(binaryStr); i=i+4 {
        hexStr += binaryToHex(binaryStr[i:i+4])
    }
    leadingZeroCount := 0
    for i:=0; i<len(hexStr); i++ {
        if string(hexStr[i]) != "0" {
            break
        }
        leadingZeroCount++
    }
    return hexStr[leadingZeroCount:len(hexStr)]
}

func normalizeNum(a int) string {
    aBinaryStr := strconv.FormatInt(int64(a), 2)
    if a < 0 {
        aBinaryStr = aBinaryStr[1:len(aBinaryStr)]
    }
    length := len(aBinaryStr)
    for i:=0; i<32-length; i++ {
        aBinaryStr = "0" + aBinaryStr
    }
    if a < 0 {
        aBinaryStr = getComplement(aBinaryStr)
    }
    return aBinaryStr
}

func getComplement(binaryStr string) string {
    binaryStrReverse := ""
    for i:=0; i<32; i++ {
        if string(binaryStr[i]) == "1" {
            binaryStrReverse = binaryStrReverse + "0"
        } else {
            binaryStrReverse = binaryStrReverse + "1"
        }
    }
    return sumBinary(binaryStrReverse, normalizeNum(1))
}

func sumBinary(a string, b string) string {
    sum := ""
    sign := int64(0)
    for i:=0; i<32; i++ {
        aPos, _ := strconv.ParseInt(string(a[31-i]), 10, 32)
        bPos, _ := strconv.ParseInt(string(b[31-i]), 10, 32)
        sum = strconv.FormatInt(int64(aPos ^ bPos ^ sign), 10) + sum
        if aPos == int64(1) && (bPos == int64(1) || sign == int64(1)) {
            sign = 1
        } else if bPos == int64(1) && sign == int64(1) {
            sign = 1
        } else {
            sign = 0
        }
    }
    return sum
}

func binaryToHex(str string) string {
    hex := ""
    sum := 0
    for i:=0; i<len(str); i++ {
        if string(str[len(str)-1-i]) == "1" {
            sum = sum + int(math.Pow(2, float64(i)))
        }
    }
    if sum == 0 {
        return "0"
    }
    if sum >= 10 {
        hex = string('a' + sum - 10)
        return hex
    }
    hex = strconv.Itoa(sum)
    return hex
}