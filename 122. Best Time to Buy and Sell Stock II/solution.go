func maxProfit(prices []int) int {
    maxProfit := 0
    if len(prices) == 0 {
        return maxProfit
    }
    cur := prices[0]
    for i:=1; i<len(prices); i++ {
        if prices[i] > cur {
            if i + 1 >= len(prices) {
                maxProfit = maxProfit + (prices[i] - cur)   
                break
            }
            if prices[i+1] > prices[i] {
                continue
            } else {
                maxProfit = maxProfit + (prices[i] - cur)   
                cur = prices[i+1]
                i++
            }
        } else {
            cur = prices[i]
        }
    }
    return maxProfit
}