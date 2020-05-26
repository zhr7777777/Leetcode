/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    if(prices.length < 2) return 0
    let curMin = prices[0]
    let profit = 0
    for(let i=1; i<prices.length; i++) {
        if(prices[i] - curMin > profit) {
            profit = prices[i] - curMin
        }
        if(prices[i] < curMin) {
            curMin = prices[i]
        }
    }
    return profit
};