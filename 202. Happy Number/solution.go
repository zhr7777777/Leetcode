func isHappy(n int) bool {
    result := true
    mapping := map[int]int{}
    for {
        if n == 1 {
            break
        }
        _, has := mapping[n]
        if has == true {
            result = false
            break
        }
        mapping[n] = n
        n = cal(n)
    }
    return result
}

func cal(n int) int {
    sum := 0
    for {
        if n == 0 {
            break
        }
        remainder := (n % 10) 
        sum = sum + remainder * remainder
        n = int(n / 10)
    }
    return sum
}