package raindrops

import "strconv"

const testVersion = 3

func Convert(number int) string {
	if number%3 != 0 && number%5 != 0 && number%7 != 0 {
		return strconv.Itoa(number)
	}
	sounds := ""
	if number%3 == 0 {
		sounds += "Pling"
	}
	if number%5 == 0 {
		sounds += "Plang"
	}
	if number%7 == 0 {
		sounds += "Plong"
	}
	return sounds
}