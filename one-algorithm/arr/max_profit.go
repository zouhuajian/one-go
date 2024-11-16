package arr

import "math"

func getMax(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/?envType=study-plan-v2&envId=top-interview-150
// 121. 买卖股票的最佳时机
// 给定一个数组 prices ，它的第 i 个元素 prices[i] 表示一支给定股票第 i 天的价格。
// 你只能选择 某一天 买入这只股票，并选择在 未来的某一个不同的日子 卖出该股票。设计一个算法来计算你所能获取的最大利润。
// 返回你可以从这笔交易中获取的最大利润。如果你不能获取任何利润，返回 0 。
func maxProfit1_1(prices []int) int {
	profit := 0
	minPrice := math.MaxInt32
	for i := range prices {
		price := prices[i]
		if price < minPrice {
			minPrice = price
		} else if price-minPrice > profit {
			profit = price - minPrice
		}
	}

	return profit
}

// 使用动态规划实现
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/solutions/38477/bao-li-mei-ju-dong-tai-gui-hua-chai-fen-si-xiang-b/
// https://blog.csdn.net/lw_power/article/details/103772951
func maxProfit1_2(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}

	dp := make([][2]int, l)
	// dp[i][0] 下标为 i 这天结束的时候，不持股，手上拥有的现金数
	// dp[i][1] 下标为 i 这天结束的时候，持股，手上拥有的现金数
	dp[0][0] = 0
	dp[0][1] = -prices[0]
	for i := 1; i < l; i++ {
		// 现金, 1. 保持现金；2. 卖出股票
		dp[i][0] = getMax(dp[i-1][0], dp[i-1][1]+prices[i])
		// 持股，1. 保持持股；2. 买入股票
		dp[i][1] = getMax(dp[i-1][1], -prices[i])
	}
	return dp[l-1][0]
}

func maxProfit1_3(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	cash := 0
	hold := -prices[0]
	for i := 1; i < l; i++ {
		cash = getMax(cash, hold+prices[i])
		hold = getMax(hold, -prices[i])
	}
	return cash
}

// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/description/?envType=study-plan-v2&envId=top-interview-150
// 122. 买卖股票的最佳时机 II
// 给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
// 在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
// 返回 你能获得的 最大 利润 。
// 二维状态数组/两个一维数组
func maxProfit2_1(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	// 现金
	cash := make([]int, l)
	// 持有股票
	hold := make([]int, l)
	// 第一天持有股票
	hold[0] = -prices[0]
	for i := 1; i < l; i++ {
		// 当天为现金：1.现金不买入; 2.昨天持有股票卖掉
		cash[i] = getMax(cash[i-1], hold[i-1]+prices[i])
		// 当天持有股票：1.继续持有股票; 2.现金买入股票
		hold[i] = getMax(hold[i-1], -prices[i]+cash[i-1])
	}
	return cash[l-1]
}

// 两个常量
func maxProfit2_2(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	// 现金
	cash := make([]int, l)
	// 持有股票
	hold := make([]int, l)
	// 第一天持有股票
	hold[0] = -prices[0]
	for i := 1; i < l; i++ {
		price := prices[i]
		// 当天持有股票：1.继续持有股票; 2.现金买入股票
		hold[i] = getMax(hold[i-1], cash[i-1]-price)
		// 当天为现金：1.现金不买入; 2.昨天持有股票卖掉
		cash[i] = getMax(cash[i-1], hold[i-1]+price)
	}
	return cash[l-1]
}

// 贪心算法
func maxProfit2_3(prices []int) int {
	l := len(prices)
	if l < 2 {
		return 0
	}
	sum := 0
	for i := 1; i < l; i++ {
		if prices[i-1] < prices[i] {
			sum += prices[i] - prices[i-1]
		}
	}
	return sum
}

// 123. 买卖股票的最佳时机 III
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/description/
// 给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
// 4个一维数组
func maxProfit3_1(prices []int) int {
	l := len(prices)
	// 啥也不干

	// 第一次持有
	hold1 := make([]int, l)
	// 第一次卖出
	cash1 := make([]int, l)
	// 第二次持有
	hold2 := make([]int, l)
	cash2 := make([]int, l)
	hold1[0] = -prices[0]
	hold2[0] = -math.MaxInt32
	for i := 1; i < l; i++ {
		price := prices[i]
		hold1[i] = getMax(hold1[i-1], 0-price)
		cash1[i] = getMax(cash1[i-1], hold1[i-1]+price)
		hold2[i] = getMax(hold2[i-1], cash1[i-1]-price)
		cash2[i] = getMax(cash2[i-1], hold2[i-1]+price)
	}
	return getMax(cash1[l-1], cash2[l-1])
}

// 309. 买卖股票的最佳时机含冷冻期
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/description/
// 给定一个整数数组prices，其中第  prices[i] 表示第 i 天的股票价格 。​
// 设计一个算法计算出最大利润。在满足以下约束条件下，你可以尽可能地完成更多的交易（多次买卖一支股票）:
// 卖出股票后，你无法在第二天买入股票 (即冷冻期为 1 天)。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfit4_1(prices []int) int {
	l := len(prices)
	// 无动作
	// 买入
	hold := make([]int, l)
	// 卖出
	cash := make([]int, l)
	// 冻结期
	pause := make([]int, l)
	hold[0] = -prices[0]
	for i := 1; i < l; i++ {
		price := prices[i]
		// 不持有且不操作，冷冻期不持有
		cash[i] = getMax(cash[i-1], pause[i-1])
		// 持有但不操作，新买入
		hold[i] = getMax(hold[i-1], -price+cash[i-1])
		// 卖出并进入冷冻期
		pause[i] = hold[i-1] + price
	}
	return getMax(cash[l-1], pause[l-1])
}

// 188. 买卖股票的最佳时机 IV
// https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/description/
// 给你一个整数数组 prices 和一个整数 k ，其中 prices[i] 是某支给定的股票在第 i 天的价格。
// 设计一个算法来计算你所能获取的最大利润。你最多可以完成 k 笔交易。也就是说，你最多可以买 k 次，卖 k 次。
// 注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
func maxProfit5_1(k int, prices []int) int {

	return 0
}
