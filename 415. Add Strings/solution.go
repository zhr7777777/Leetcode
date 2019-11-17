func addStrings(num1 string, num2 string) string {
    sign := 0
    result := ""
    if len(num1) > len(num2) {
        num2 = addLeadingZero(num2, len(num1) - len(num2))
    } else if len(num1) < len(num2) {
        num1 = addLeadingZero(num1, len(num2) - len(num1))
    }
    
    for i:=0; i<len(num1); i++ {
        l1, _ := strconv.Atoi(string(num1[len(num1)-1-i]))
        l2, _ := strconv.Atoi(string(num2[len(num2)-1-i]))
        sum := l1 + l2 + sign
        l := strconv.Itoa(sum)
        if sum >= 10 {
            sign = 1
            l = strconv.Itoa(sum % 10)
        } else {
            sign = 0
        }
        result = l + result
    }
    if sign == 1 {
        result = "1" + result
    }
    return result
}

func addLeadingZero(str string, num int) string {
    for i:=0; i<num; i++ {
        str = "0" + str
    }
    return str
}