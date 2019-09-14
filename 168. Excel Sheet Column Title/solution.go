func convertToTitle(n int) string {
    if n <= 26 {
        return string(n + 64)
    }
    cur := n
    title := ""
    for {
        remainder := cur % 26
        quotient := int(cur / 26)
        if remainder == 0 { // important
            remainder = 26
            quotient--
        }
        if quotient < 26 {
            title = string(remainder + 64) + title
            if quotient > 0 {   // important
                title = string(quotient + 64) + title   
            }
            break
        } else {
            title = string(remainder + 64) + title
            cur = quotient
        }
    }
    return title
}