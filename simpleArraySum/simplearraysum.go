package simpleArraySum

//Given an array of integers, find the sum of its elements.

func simpleArraySum(ar []int32) int32 {
	var sum int32

	for _, a := range ar {
		sum += a
	}

	return sum
}
