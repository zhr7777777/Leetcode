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

func countPrimes(n int) int {       // fast solution
    isPrime := make([]bool, n)
    // isPrime := [n]bool{}     //  non-constant array bound n
    for i:=2; i<n; i++ {
        isPrime[i] = true
    }
    for i:=2; i*i<n; i++ {
        if !isPrime[i] {
            continue
        }
        for j:=i*i; j<n; j+=i {
            isPrime[j] = false
        }
    }
    count :=0
    for i:=2; i<n; i++ {
        if isPrime[i] {
            count++
        }
    }
    return count
}