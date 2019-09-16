func trailingZeroes(n int) int {
    pow := log5(n)
    result := 0
    for i:=0; i<pow; i++ {  // important
        result += int(n / pow5(i+1))
    }
    return result
}

func log5(n int) int {
    count := 0
    cur := 5
    for {
        if cur > n {
            break
        } else {
            count++
        }
        cur *= 5
    }
    return count
}

func pow5(n int) int {
    result := 1
    for i:=0; i<n; i++ {
        result *= 5
    }
    return result
}