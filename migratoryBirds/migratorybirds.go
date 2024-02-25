package migratoryBirds

/*
 * Complete the 'migratoryBirds' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts INTEGER_ARRAY arr as parameter.
 */

func migratoryBirds(arr []int32) int32 {
	birdsMap := make(map[int32]int32, len(arr))
	for _, bird := range arr {
		birdsMap[bird]++
	}
	var maxBird int32
	var maxCount int32
	for bird, count := range birdsMap {
		if count > maxCount {
			maxCount = count
			maxBird = bird
		} else if count == maxCount && bird < maxBird {
			maxBird = bird
		}
	}
	return maxBird
}
