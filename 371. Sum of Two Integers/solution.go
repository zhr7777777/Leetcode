func getSum(a int, b int) int {
    strA := normalizeNum(a)
    strB := normalizeNum(b)
    sum := sumBinary(strA, strB)
    if string(sum[0]) == "1" {
        sum = "-" + getComplement(sum)
    }
    s, _ := strconv.ParseInt(sum, 2, 32)
    return int(s)
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