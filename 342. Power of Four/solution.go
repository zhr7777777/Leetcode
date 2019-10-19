func isPowerOfFour(num int) bool {
    epsilon := 0.0000000001
    _, decimal := math.Modf(math.Log2(float64(num)) / 2)
    return decimal + epsilon <= 2 * epsilon
}