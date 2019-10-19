func isPowerOfThree(n int) bool {
    epsilon := 0.0000000001
    _, decimal := math.Modf(math.Log(float64(n)) / math.Log(float64(3)))
    return decimal + epsilon <= 2 * epsilon
}