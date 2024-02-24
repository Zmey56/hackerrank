package breakingRecords

/*
 * Complete the 'breakingRecords' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts INTEGER_ARRAY scores as parameter.
 */

func breakingRecords(scores []int32) []int32 {
	// Write your code here
	var max int32 = scores[0]
	var min int32 = scores[0]
	var maxCount int32
	var minCount int32
	for _, score := range scores {
		if score > max {
			max = score
			maxCount++
		}
		if score < min {
			min = score
			minCount++
		}
	}
	return []int32{maxCount, minCount}

}
