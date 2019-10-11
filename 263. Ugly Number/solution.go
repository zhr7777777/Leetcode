func isUgly(num int) bool {
    if num <= 0 {
        return false
    }
    for {
        if num == 1 {
            break
        }
        count := 0
        fs := []int{2, 3, 5}
        for i:=0; i<len(fs); i++ {
            if num % fs[i] == 0 {
                num = num / fs[i]
            } else {
                count++
            }
        }
        if count == 3 {
            return false
        }
    }
    return true
}