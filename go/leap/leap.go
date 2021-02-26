// Package leap implements a method to calculate whether a given year is a leap year.
package leap

// IsLeapYear - takes an int `year` and return a bool describing whether the given year is a leap year or not
func IsLeapYear(year int) bool {
	isLeap := false
	if year%4 == 0 {
		if year%100 == 0 {
			if year%400 == 0 {
				isLeap = true
			} else {
				isLeap = false
			}
		} else {
			isLeap = true
		}
	} else {
		isLeap = false
	}

	return isLeap
}
