func canPlaceFlowers(flowerbed []int, n int) bool {
    zeroCount := 0
    canPlaceCount := 0
    hasFlower := false
    flowerbed = append([]int{0}, flowerbed...)
    flowerbed = append(flowerbed, 0)
    for i:=0; i<len(flowerbed); i++ {
        if flowerbed[i] == 0 {
            zeroCount = zeroCount + 1
            if i == len(flowerbed) - 1 {    // should consider when loop is finished
                canPlaceCount = canPlaceCount + int((zeroCount - 1) / 2)
            }
        } else {
            hasFlower = true
            canPlaceCount = canPlaceCount + int((zeroCount - 1) / 2)
            zeroCount = 0
        }
    }
    if !hasFlower {
        return int((len(flowerbed) - 2 + 1) / 2) >= n // should deduct 2
    }
    return canPlaceCount >= n
}