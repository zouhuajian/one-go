package main

// 合并两个有序数组
// https://leetcode.cn/problems/merge-sorted-array/
// https://leetcode.cn/problems/merge-sorted-array/solutions/14030/si-xiang-mei-you-chuang-xin-de-di-fang-zhu-yao-ti-/
func merge(nums1 []int, m int, nums2 []int, n int) {
	index1 := m - 1
	index2 := n - 1
	index := m + n - 1
	for i := index; i >= 0; i-- {
		if index1 < 0 {
			copy(nums1, nums2[:index2+1])
			return
		} else if index2 < 0 {
			return
		} else if nums1[index1] >= nums2[index2] {
			nums1[i] = nums1[index1]
			index1--
		} else {
			nums1[i] = nums2[index2]
			index2--
		}
	}
}

/*func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	//nums1 := []int{0}
	//nums2 := []int{1}
	merge(nums1, 3, nums2, 3)
	fmt.Printf("nums1: %+v", nums1)
}
*/
