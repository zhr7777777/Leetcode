func reverseBits(num uint32) uint32 {
    var front, back, f, b uint32
    front, back = 1 << 31, 1
    for i := 0; i < 16; i++ {
        f = front & num
        b = back & num
        if f == 0 && b != 0 {
            num = num | front
            num = num ^ back
        } else if f != 0 && b == 0 {
            num = num ^ front
            num = num | back
        }
        front, back = front >> 1, back << 1
    }
    return num
}