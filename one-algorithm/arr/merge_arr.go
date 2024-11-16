package arr

// https://leetcode.cn/problems/merge-sorted-array/description/?envType=study-plan-v2&envId=top-interview-150
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	index := m + n - 1
	var currentVal int
	for index1 >= 0 || index2 >= 0 {
		if index1 < 0 {
			currentVal = nums2[index2]
			index2--
		} else if index2 < 0 {
			currentVal = nums1[index1]
			index1--
		} else if nums1[index1] >= nums2[index2] {
			currentVal = nums1[index1]
			index1--
		} else {
			currentVal = nums2[index2]
			index2--
		}
		nums1[index] = currentVal
		index--
	}
}
