package staircase

import (
	"fmt"
	"strings"
)

//Staircase detail
//This is a staircase of size n = 4:
//   #
//  ##
// ###
//####
//Its base and height are both equal to n. It is drawn using # symbols and spaces. The last line is not preceded by any spaces.
//Write a program that prints a staircase of size .

func staircase(n int32) {
	symbol := "#"
	space := " "
	result := ""

	for i := int32(1); i <= int32(n); i++ {
		result = strings.Repeat(space, int(n)-int(i)) + strings.Repeat(symbol, int(i))
		fmt.Println(result)
	}

}
