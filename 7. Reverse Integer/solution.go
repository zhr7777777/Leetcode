func reverse(x int) int {
    var result int = 0
    for {
        if(x == 0) {
            break
        }
        var cur int = x % 10
        x = int(x / 10)
        result = result * 10 + cur
    }
    if(result >2147483647 || result < -2147483648) {
        return 0
    }
    return result
}