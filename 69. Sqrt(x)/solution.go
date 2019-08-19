func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    if x == 1 {
        return 1
    }
    result := 0
    for i:=1; i<x; i++ {
        if i * i <= x && (i+1) * (i+1) > x {
            result = i
            break
        }
    }
    return result
}