package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 0}
	sortArray(nums)
	fmt.Println(nums)
}

// 归并排序
func sortArray(nums []int) []int {
	tmp := make([]int, len(nums))
	mergeSort(nums, tmp, 0, len(nums)-1)
	return nums
}

func mergeSort(nums []int, tmp []int, left int, right int) {
	if left == right {
		return
	}
	mid := (left + right) / 2
	mergeSort(nums, tmp, left, mid)
	mergeSort(nums, tmp, mid+1, right)
	mergeSubs(nums, tmp, left, mid, right)
}

func mergeSubs(nums []int, tmp []int, left, mid, right int) {
	leftIndex := left
	rightIndex := mid + 1
	tmpIndex := left

	for leftIndex <= mid && rightIndex <= right {
		if nums[leftIndex] <= nums[rightIndex] {
			tmp[tmpIndex] = nums[leftIndex]
			leftIndex++
		} else {
			tmp[tmpIndex] = nums[rightIndex]
			rightIndex++
		}
		tmpIndex++
	}

	for leftIndex <= mid {
		tmp[tmpIndex] = nums[leftIndex]
		leftIndex++
		tmpIndex++
	}

	for rightIndex <= right {
		tmp[tmpIndex] = nums[rightIndex]
		rightIndex++
		tmpIndex++
	}

	for i := left; i <= right; i++ {
		nums[i] = tmp[i]
	}
}
