package arr

// https://leetcode.cn/problems/majority-element/?envType=study-plan-v2&envId=top-interview-150
// 169. 多数元素
// 给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
// 2024-02-26
func majorityElement(nums []int) int {
	count := 0
	var result int
	for _, num := range nums {
		if count == 0 {
			result = num
		}
		if result == num {
			count++
		} else {
			count--
		}
	}
	return result
}

func majorityElement2(nums []int) int {
	m := make(map[int]int)
	l := len(nums)
	half := l / 2
	for i := 0; i < l; i++ {
		c := m[nums[i]]
		c++
		if c > half {
			return nums[i]
		}
		m[nums[i]] = c
	}
	return -1
}
