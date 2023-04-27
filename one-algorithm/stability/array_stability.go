package main

import (
	"fmt"
	"github.com/shopspring/decimal"
	"math"
	"sort"
)

func main() {

	// arr := []int{10, 20, 10, 10, 10, 11, 12, 12, 123, 30}
	arr := []int{21, 22, 23, 2, 20, 20, 20, 20, 20, 90}
	fmt.Println("Original array:", arr)
	filteredArr := removeOutliersViaQuartiles(arr, 2)
	// Compute stability metric (mean absolute deviation)
	stabilityScore := computeStabilityScore(filteredArr)

	fmt.Printf("Stability Score: %f\n", stabilityScore)
	// 将稳定性阈值定义为标准差的分数，当前为标准差的 n 倍
	stabilityThreshold := decimal.NewFromFloat(stabilityScore).Mul(decimal.NewFromFloat(1.0)).InexactFloat64()
	result := isStable(filteredArr, stabilityThreshold)
	fmt.Printf("Stability Result: %v\n", result)
}

func isStable(arr []int, stabilityThreshold float64) bool {
	n := len(arr)
	mean := float64(sum(arr)) / float64(n)
	notStabilityCount := 0
	for _, val := range arr {
		if math.Abs(float64(val)-mean) > stabilityThreshold {
			notStabilityCount++
		}
	}
	fmt.Printf("Stability threshold is %v, mean is %v, not stability element count: %v/%v\n",
		stabilityThreshold, mean, notStabilityCount, n)
	// 1/4 的元素低于阈值，则判断这个数组较为稳定
	return notStabilityCount <= n/4
}

// Computes the mean absolute deviation of an integer array
// 基于标准偏差计算稳定性
func computeStabilityScore(arr []int) float64 {
	n := len(arr)
	if n == 0 {
		return 0
	}
	mean := float64(sum(arr)) / float64(n)
	sumAbsDev := 0.0
	for _, val := range arr {
		sumAbsDev += math.Pow(float64(val)-mean, 2)
	}
	return math.Sqrt(sumAbsDev / float64(n))
}

// 基于IQR(四分位距)移除离群值
// threshold: 用于计算上下限阈值，比如2.0倍的
func removeOutliersViaQuartiles(arr []int, rate float64) []int {
	q1, q3 := quartiles(arr)
	iqr := q3 - q1
	lowerBound := q1 - rate*iqr
	upperBound := q3 + rate*iqr
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
