package diagonalDifference

// Given a square matrix, calculate the absolute difference between the sums of its diagonals.

func diagonalDifference(arr [][]int32) int32 {
	var sum1 int32
	var sum2 int32

	for i := 0; i < len(arr); i++ {
		sum1 += arr[i][i]
		sum2 += arr[i][len(arr)-i-1]
	}

	if sum1 >= sum2 {
		return sum1 - sum2
	} else {
		return sum2 - sum1
	}
}
