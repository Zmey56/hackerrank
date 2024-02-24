package getTotalX

/*
 * Complete the 'getTotalX' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY a
 *  2. INTEGER_ARRAY b
 */

func getTotalX(a []int32, b []int32) int32 {

	var count int32
	for i := a[len(a)-1]; i <= b[0]; i++ {
		var flag bool = true
		for _, val := range a {
			if i%val != 0 {
				flag = false
				break
			}
		}
		if flag {
			for _, val := range b {
				if val%i != 0 {
					flag = false
					break
				}
			}
		}
		if flag {
			count++
		}
	}
	return count

}
