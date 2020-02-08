func checkRecord(s string) bool {
    absentCount := 0
    lateCount := 0
    for i:=0; i<len(s); i++ {
        if i != 0 && s[i-1] != s[i] {
            lateCount = 0
        }
        if s[i] == 'A' {
            absentCount++
        } else if s[i] == 'L' {
            lateCount++
        }
        if absentCount >= 2 {
            return false
        }
        if lateCount >= 3 {
            return false
        }
    }
    return true
}