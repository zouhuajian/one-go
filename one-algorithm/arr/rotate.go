package arr

// https://leetcode.cn/problems/rotate-array/?envType=study-plan-v2&envId=top-interview-150
// 给定一个整数数组 nums，将数组中的元素向右轮转 k 个位置，其中 k 是非负数。

// Note: 三次反转，时间复杂度 O(2n) -> O(n)
func rotate(nums []int, k int) {
	l := len(nums)
	reverse(nums, 0, l-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, l-1)
}

// 带范围的数组反转
func reverse(nums []int, left int, right int) {
	for left < right {
		tmp := nums[left]
		nums[left] = nums[right]
		nums[right] = tmp
		left++
		right--
	}
}

// 多次遍历
func rotate2(nums []int, k int) {
	l := len(nums)
	for i := 0; i < k; i++ {
		tmp := nums[l-1]
		for i := l - 2; i >= 0; i-- {
			nums[i+1] = nums[i]
		}
		nums[0] = tmp
	}
}
