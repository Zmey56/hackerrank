package timeConversion

//Given a time in -hour AM/PM format, convert it to military (24-hour) time.
//Note: - 12:00:00AM on a 12-hour clock is 00:00:00 on a 24-hour clock.
//- 12:00:00PM on a 12-hour clock is 12:00:00 on a 24-hour clock.

func timeConversion(s string) string {
	result := ""
	if s[8:] == "AM" {
		if s[:2] == "12" {
			result = "00" + s[2:8]
		} else {
			result = s[:8]
		}
	} else {
		if s[:2] == "12" {
			result = s[:8]
		} else {
			result = string(rune(s[0])+1) + string(rune(s[1])+2) + s[2:8]
		}
	}
	return result
}
