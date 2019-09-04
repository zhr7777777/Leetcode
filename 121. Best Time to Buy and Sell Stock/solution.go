func maxProfit(prices []int) int {
    if len(prices) == 0 {
        return 0
    }
    max := 0
    cur := prices[0]
    for i:=1; i<len(prices); i++ {
        if prices[i] > cur {
            if prices[i] - cur > max {
                max = prices[i] - cur
            }
        } else {
            cur = prices[i]
        }
    }
    return max
}