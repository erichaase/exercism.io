package leap

const testVersion = 3

func IsLeapYear(y int) bool {
	if y % 400 == 0 {
		return true
	} else if y % 100 == 0 {
		return false
	} else if y % 4 == 0 {
		return true
	} else {
		return false
	}
}
