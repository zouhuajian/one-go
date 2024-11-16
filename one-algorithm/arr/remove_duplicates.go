package arr

// 26. 删除有序数组中的重复项
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array/?envType=study-plan-v2&envId=top-interview-150

// 输入：nums = [0,0,0,1,1,1,2,2,3,3,4]
// 输出：5, nums = [0,1,2,3,4]

func removeDuplicates1(nums []int) int {
	l := len(nums)
	if l <= 1 {
		return l
	}
	cur := 0
	for i := 0; i < l; i++ {
		if nums[cur] != nums[i] {
			cur++
			nums[cur] = nums[i]
		}
	}
	return cur + 1
}

// 80. 删除有序数组中的重复项 II
// https://leetcode.cn/problems/remove-duplicates-from-sorted-array-ii/description/?envType=study-plan-v2&envId=top-interview-150
// 输入：nums = [1,1,1,1,1,1,2,3,3]
// 输出：7, nums = [0,0,1,1,2,3,3]
// 使用快慢指针
func removeDuplicates(nums []int) int {
	l := len(nums)
	if l <= 2 {
		return l
	}
	slow := 2
	fast := 2
	for fast < l {
		if nums[slow-2] != nums[fast] {
			nums[slow] = nums[fast]
			slow++
		}
		fast++
	}
	return slow
}
