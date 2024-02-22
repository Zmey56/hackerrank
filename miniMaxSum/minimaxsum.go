package miniMaxSum

import (
	"fmt"
)

//Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly
//four of the five integers. Then print the respective minimum and maximum values as a single line
//of two space-separated long integers.

func miniMaxSum(arr []int32) [2]int {
	var min int
	var max int
	var minSum int
	var maxSum int
	var sum int
	for i, a := range arr {
		a := int(a)
		if i == 0 {
			min = a
			max = a
		}
		if a < min {
			min = a
		}
		if a > max {
			max = a
		}
	}
	for _, a := range arr {
		a := int(a)
		sum += a
	}
	minSum = sum - max
	maxSum = sum - min
	fmt.Println(minSum, maxSum)

	return [2]int{minSum, maxSum}
}
