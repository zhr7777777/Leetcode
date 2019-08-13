func countAndSay(n int) string {
    str := "1"
    for i:=1; i<n; i++ {
        num := 1
        cur := str
        str = ""
        for j:=0; j<len(cur); j++ {
            if j == len(cur) - 1 {
                str = str + fmt.Sprintf("%d", num) + string(cur[j])
                break
            }
            if string(cur[j]) == string(cur[j+1]) {
                num++
            } else {
                str = str + fmt.Sprintf("%d", num)  + string(cur[j])
                num = 1
            }
        }
    }
    return str
}