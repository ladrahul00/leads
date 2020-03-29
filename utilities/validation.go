package utilities

import "regexp"

// ValidateString if null or empty
func ValidateString(value string) bool {
	if len(value) == 0 {
		return false
	}
	return true
}

// ValidateEmail pop
func ValidateEmail(value string) bool {
	emailRegexp := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	return emailRegexp.MatchString(value)
}
