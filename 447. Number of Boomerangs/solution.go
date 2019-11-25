func numberOfBoomerangs(points [][]int) int {
    num := 0
    for i:=0; i<len(points); i++ {
        cur := points[i]
        mapping := map[int]int{}
        for j:=0; j<len(points); j++ {
            if j != i {
                distanceSquare := getDistanceSquare(cur, points[j])
                count, has := mapping[distanceSquare]
                if has {
                    mapping[distanceSquare] = count + 1
                } else {
                    mapping[distanceSquare] = 1
                }   
            }
        }
        for distance := range mapping {
            count := mapping[distance]
            if count != 1 {
                num = num + getSum(count) * 2
            }
        }
    }
    return num
}

func getDistanceSquare(a []int, b []int) int {
    x := a[0] - b[0]
    y := a[1] - b[1]
    return x * x + y * y
}

func getSum(n int) int {
    if n == 0 {
        return 0
    }
    return n - 1 + getSum(n-1) 
}