var result []string

func fizzBuzz(n int) []string {
    if result != nil && len(result) >= n {
        return result[:n]
    }
    start := 0
    if result != nil {
        start = len(result)
    }
    for i:=start; i<n; i++ {
        num := i + 1
        if num % 3 == 0 && num % 5 == 0 {
            result = append(result, "FizzBuzz")
            continue
        }
        if num % 3 == 0 {
            result = append(result, "Fizz")
            continue
        }
        if num % 5 == 0 {
            result = append(result, "Buzz")
            continue
        }
        result = append(result, strconv.Itoa(num))
    }
    return result
}