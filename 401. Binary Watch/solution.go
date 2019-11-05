func readBinaryWatch(num int) []string {
    all := []string{}
    for i:=0; i<12; i++ {
        for j:=0; j<60; j++ {
            if hammingWeight(i) + hammingWeight(j) == num {
                hour := strconv.Itoa(i)
                min := strconv.Itoa(j)
                if j < 10 {
                    min = "0" + min
                }
                all = append(all, hour + ":" + min)
            }
        }
    }
    return all
}

func hammingWeight(num int) int {
    count := 0
    for i:=0; i<6; i++ {
        if num & 1 == 1 {
            count++
        }
        num = num >> 1
    }
    return count
}