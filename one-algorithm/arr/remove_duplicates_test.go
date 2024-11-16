package arr

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	l := removeDuplicates(nums)
	println(l)
	for i := 0; i < l; i++ {
		print(nums[i], ",")
	}
}
