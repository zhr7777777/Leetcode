func findComplement(num int) int {
    complementStr := ""
    for {
        if num == 0 {
            break
        }
        if num & 1 == 1 {
            complementStr = "0" + complementStr
        } else {
            complementStr = "1" + complementStr
        }
        num = num >> 1
    }
    complement, _ := strconv.ParseInt(complementStr, 2, 32)
    return int(complement)
}