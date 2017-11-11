package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	/*
		on every year that is evenly divisible by 4
			except every year that is evenly divisible by 100
		    unless the year is also evenly divisible by 400
	*/

	if year%4 == 0 {
		if year%400 == 0 {
			return true
		} else {
			if year%100 == 0 {
				return false
			}
			return true
		}
	} else {
		return false
	}

}
