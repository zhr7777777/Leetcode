func checkPerfectNumber(num int) bool {
    if num == 1 {
        return false
    }
    sum := 1
    n := 2
    for {
        if n > int(math.Sqrt(float64(num))) {
            break
        }
        if sum > num {
            return false
        }
        if num % n == 0 {
            sum += n
            sum = sum + num / n
        }
        n++
    }
    if sum == num {
        return true
    }
    return false
}