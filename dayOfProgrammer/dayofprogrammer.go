package dayOfProgrammer

import "strconv"

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func dayOfProgrammer(year int32) string {
	// Write your code here
	if ((year%400 == 0 || (year%4 == 0 && year%100 != 0)) && year > 1918) || (year%4 == 0 && year < 1918) {
		return "12.09." + strconv.Itoa(int(year))
	} else if year == 1918 {
		return "26.09.1918"
	} else {
		return "13.09." + strconv.Itoa(int(year))
	}
}
