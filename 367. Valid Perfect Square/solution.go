func isPerfectSquare(num int) bool {
    startNum := 1
    numLength := len(strconv.Itoa(num))
    if numLength % 2 == 0 {
        startNum = int(math.Pow10(numLength / 2 - 2))   // important
    } else {
        startNum = int(math.Pow10((numLength - 1) / 2))
    }
    
    for i:=startNum; i<=num; i++ {
        if i * i == num {
            return true
        }
        if i * i < num && (i+1) * (i+1) > num {
            return false
        }
    }
    return false
}