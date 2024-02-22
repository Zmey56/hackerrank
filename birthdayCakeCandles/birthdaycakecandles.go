package birthdayCakeCandles

//You are in charge of the cake for a child's birthday.
//You have decided the cake will have one candle for each year of their total age.
//They will only be able to blow out the tallest of the candles. Count how many candles are tallest.

func birthdayCakeCandles(candles []int32) int32 {
	var max int32
	var count int32

	for _, c := range candles {
		if c > max {
			max = c
			count = 1
		} else if c == max {
			count++
		}
	}

	return count

}
