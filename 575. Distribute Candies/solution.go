func distributeCandies(candies []int) int {
    mapping := map[int]int{}
    count := 0
    for i:=0; i<len(candies); i++ {
        _, has := mapping[candies[i]]
        if !has {
            count++
            mapping[candies[i]] = 0
        }
    }
    if count >= len(candies) / 2 {
        return len(candies) / 2
    }
    return count
}