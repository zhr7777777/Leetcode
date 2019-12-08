func hammingDistance(x int, y int) int {
    dis := 0
    for i:=0; i<32; i++ {
        if (x & 1) ^ (y & 1) == 1 {
            dis++
        }
        x = x >> 1
        y = y >> 1
    }
    return dis
}