package utils

import (
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"sort"
	"strings"
	"unicode"
)

const specialChars = "!@#$%^&*_+-=[]{}|<>?/~"

func Map2String(m map[string]interface{}, split string) string {
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	var msg strings.Builder
	for i, key := range keys {
		if i > 0 {
			msg.WriteString(split)
		}
		msg.WriteString(fmt.Sprintf("%s=%v", key, m[key]))
	}
	return msg.String()
}

func HashPassword(password string, salt []byte) string {
	passwordWithSalt := append([]byte(password), salt...)
	hash := sha512.Sum512(passwordWithSalt)
	hashHex := hex.EncodeToString(hash[:])
	return hashHex
}

func ValidatePassword(password string) bool {
	var (
		hasUpper   bool // 是否包含大写字母
		hasLower   bool // 是否包含小写字母
		hasDigit   bool // 是否包含数字
		hasSpecial bool // 是否包含特殊字符
	)
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		case contains(specialChars, char):
			hasSpecial = true
		default:
			return false
		}
	}

	count := 0
	if hasUpper {
		count++
	}
	if hasLower {
		count++
	}
	if hasDigit {
		count++
	}
	if hasSpecial {
		count++
	}
	return len(password) >= 8 && len(password) <= 18 && count >= 3
}

func contains(specialChars string, char rune) bool {
	for _, c := range specialChars {
		if c == char {
			return true
		}
	}
	return false
}
