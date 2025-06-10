package utils

import "regexp"

// check password is: longer than 8, uppercase, lowercase, number and symbol
func IsStrongPassword(password string) bool {
	var (
		hasMinLen  = len(password) >= 8
		hasUpper   = regexp.MustCompile(`[A-Z]`).MatchString(password)
		hasLower   = regexp.MustCompile(`[a-z]`).MatchString(password)
		hasNumber  = regexp.MustCompile(`[0-9]`).MatchString(password)
		hasSpecial = regexp.MustCompile(`[!@#~$%^&*()_+|<>?:{}]`).MatchString(password)
	)

	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}
