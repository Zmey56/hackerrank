package bonAppetit

import "fmt"

/*
 * Complete the 'bonAppetit' function below.
 *
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY bill
 *  2. INTEGER k
 *  3. INTEGER b
 */

func bonAppetit(bill []int32, k int32, b int32) {
	// Write your code here
	var sum int32
	for i, val := range bill {
		if int32(i) != k {
			sum += val
		}
	}
	sum = sum / 2
	if sum == b {
		fmt.Println("Bon Appetit")
	} else {
		fmt.Println(b - sum)
	}
}
