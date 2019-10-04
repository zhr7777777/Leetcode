func isPowerOfTwo(n int) bool {
    if n <= 0 { // important
        return false
    }
    for {
        if n == 1 {
            return true
        }
        if n % 2 == 1 {
            return false
        }
        n = n / 2
    }
    return true
}