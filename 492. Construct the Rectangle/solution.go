func constructRectangle(area int) []int {
    width := math.Sqrt(float64(area))
    w := int(width)
    rec := []int{}
    if width == float64(w) {
        rec = append(rec, w)
        rec = append(rec, w)
        return rec
    }
    for i:=w+1; i<=area; i++ {
        if area % i == 0 {
            rec = append(rec, i)
            rec = append(rec, area / i)
            break
        }
    }
    return rec
}