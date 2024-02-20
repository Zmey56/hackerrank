package plusMinus

import "fmt"

//Given an array of integers, calculate the ratios of its elements that are positive, negative, and zero.
//Print the decimal value of each fraction on a new line with places after the decimal.

func plusMinus(arr []int32) {

	var positiveCount float64
	var negativeCount float64
	var zeroCount float64
	var total float64

	for _, a := range arr {
		if a > 0 {
			positiveCount++
		} else if a < 0 {
			negativeCount++
		} else {
			zeroCount++
		}
		total++
	}

	fmt.Println(positiveCount / total)
	fmt.Println(negativeCount / total)
	fmt.Println(zeroCount / total)

}
