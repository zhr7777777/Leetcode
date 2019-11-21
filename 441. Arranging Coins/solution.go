func arrangeCoins(n int) int {
    remaining := n
    i := 0
    for; i<n; i++ {
        remaining = remaining - (i + 1)
        if remaining < 0 {
            break
        }
    }
    return i
}