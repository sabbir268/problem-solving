function findMaxProfit(prices) {
    let maxProfit = 0;
    let minProfit = prices[0];
    for (let i = 1; i < prices.length; i++) {
        if (prices[i] < minProfit) {
            minProfit = prices[i];
        } else {
            if (prices[i] - minProfit > maxProfit) {
                maxProfit = prices[i] - minProfit;
            }
        }
    }
    return maxProfit;
}

console.log(findMaxProfit([7, 6, 4, 3, 1]));

// Time Complexity: O(n)
// Space Complexity: O(1)