package arr

// 509. 斐波那契数
func fib(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func fib2(n int) int {
	if n <= 1 {
		return n
	}
	result := 0
	prePre := 0
	pre := 1
	for i := 2; i <= n; i++ {
		result = pre + prePre
		prePre = pre
		pre = result
	}
	return result
}
