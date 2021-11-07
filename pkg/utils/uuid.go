package utils

import (
	"fmt"

	"github.com/lithammer/shortuuid/v3"
)

//UUIDWithString generate uuid with ascii name for readable id {name}_{uuid}
func UUIDWithString(name string) string {
	u := shortuuid.New()
	return fmt.Sprintf("%s_%s", cleanString([]byte(name)), u)
}

func cleanString(s []byte) string {
	j := 0
	for _, b := range s {
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			s[j] = b
			j++
		}
	}
	return string(s[:j])
}
