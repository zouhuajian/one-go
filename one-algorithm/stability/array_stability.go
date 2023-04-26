package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 100, 200, 300}
	fmt.Println("Original array:", arr)

	// Compute stability metric (mean absolute deviation)
	stability := computeStabilityWithQuartiles(arr)
	fmt.Printf("Stability: %f\n", stability)
}

// Computes the mean absolute deviation of an integer array
func computeStabilityWithQuartiles(arr []int) float64 {
	// Remove outliers using IQR method
	filteredArr := removeOutliersViaQuartiles(arr, 5.0)
	n := len(filteredArr)
	if n == 0 {
		return 0
	}
	mean := float64(sum(filteredArr)) / float64(n)
	sumAbsDev := 0.0
	for _, val := range filteredArr {
		sumAbsDev += math.Abs(float64(val) - mean)
	}
	return sumAbsDev / float64(n)
}

// 基于IQR(四分位距)移除离群值
// threshold: 用于计算上下限阈值，比如2.0倍的
func removeOutliersViaQuartiles(arr []int, threshold float64) []int {
	q1, q3 := quartiles(arr)
	iqr := q3 - q1
	lowerBound := q1 - threshold*iqr
	upperBound := q3 + threshold*iqr
	filteredArr := make([]int, 0, len(arr))
	for _, val := range arr {
		if float64(val) >= lowerBound && float64(val) <= upperBound {
			filteredArr = append(filteredArr, val)
		}
	}
	fmt.Println("Array with outliers removed:", filteredArr)
	return filteredArr
}

// Computes the first and third quartiles of an integer array
func quartiles(arr []int) (float64, float64) {
	sort.Ints(arr)
	n := len(arr)
	q1Index := int(math.Floor(float64(n-1) / 4))
	q3Index := int(math.Ceil(3 * float64(n-1) / 4))
	q1 := float64(arr[q1Index])
	q3 := float64(arr[q3Index])
	return q1, q3
}

// Computes the sum of an integer array
func sum(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}
