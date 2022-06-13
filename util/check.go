package util

import "regexp"

// CheckUsername 检查username
func CheckUsername(username string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9_.@-]{4,30}$", username); !ok {
		return false
	}
	return true
}

// CheckPassword 检查password
func CheckPassword(password string) (b bool) {
	if ok, _ := regexp.MatchString("^[a-zA-Z0-9]{4,16}$", password); !ok {
		return false
	}
	return true
}
