func findRadius(houses []int, heaters []int) int {
    if len(houses) == 0 {
        return 0
    }
    dis := 0
    for i:=0; i<len(houses); i++ {
        radius := abs(heaters[0] - houses[i])
        for j:=1; j<len(heaters); j++ {
            if abs(heaters[j] - houses[i]) < radius {
                radius = abs(heaters[j] - houses[i])
            }
            if heaters[j] == houses[i] {
                radius = 0
                break
            }
        }
        if radius > dis {
            dis = radius
        }
    }
    return dis
}

func abs(n int) int {
    i := float64(n)
    return int(math.Abs(i))
}