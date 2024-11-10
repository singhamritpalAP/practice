package solutions

func MaxProfit(prices []int) int {
	maxProfit := 0
	currentBuyPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < currentBuyPrice {
			currentBuyPrice = prices[i]
		}
		if prices[i]-currentBuyPrice > maxProfit {
			maxProfit = prices[i] - currentBuyPrice
		}
	}
	return maxProfit
}

/*
[7,1,5,3,6,4]
CBP = 1; MP = 0
CBP = 1; MP = 4
CBP = 1; MP = 4
CBP = 1; MP = 5

[7,6,4,3,1]
CBP = 7; MP = 0
CBP = 6; MP = 0
CBP = 4; MP = 0
CBP = 3; MP = 0
CBP = 1; MP = 0
*/
