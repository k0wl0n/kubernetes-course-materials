package validators

import (
	"fmt"
	"regexp"
	"strings"
)

const passwordLength = 8

func containsDigit(s string) bool {
	for _, char := range s {
		if '0' <= char && char <= '9' {
			return true
		}
	}
	return false
}

func containsSpecialChar(s string) bool {
	specialCharPattern := `[!@#$%^&*()]`
	re := regexp.MustCompile(specialCharPattern)
	return re.MatchString(s)
}

// Perform basic validation checks on the given password string
func PasswordValidator(password string, email string) error {
	if strings.Contains(password, " ") {
		return fmt.Errorf("password should not contain empty spaces")
	}

	if len(password) < passwordLength {
		return fmt.Errorf("password must be at least %d characters long", passwordLength)
	}

	if strings.ToLower(password) == password {
		return fmt.Errorf("password should contain at least one upper case character")
	}

	if strings.ToUpper(password) == password {
		return fmt.Errorf("password should contain at least one lower case character")
	}

	if !containsDigit(password) {
		return fmt.Errorf("password should contain at least one digit")
	}

	if !containsSpecialChar(password) {
		return fmt.Errorf("password should contain at least one special character")
	}

	if strings.Contains(password, email) {
		return fmt.Errorf("password should not contain your name or email address")
	}

	return nil
}
