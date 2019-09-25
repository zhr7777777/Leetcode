func countPrimes(n int) int {
    if n <= 2 {
        return 0
    }
    count := 1
    primes := []int{}
    for i:=3; i<n; i++ {
        if isPrime(i, primes) {
            primes = append(primes, i)
            count++
        }
    }
    return count
}

func isPrime(n int, primes []int) bool {
    if n == 2 {
        return true
    }
    if n % 2 == 0 {
        return false
    }
    for i:=0; i<len(primes); i++ {
        if n % primes[i] == 0 {
            return false
        }
    }
    // for i:=3; i<=int(n/2); i++ {
    //     if n % i == 0 {
    //         result = false
    //         break
    //     }
    // }
    return true
}