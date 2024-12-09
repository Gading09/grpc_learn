package pkg

import (
	"os"
	"strings"
)

func Getenv(key string) (fallback string) {
	var value string
	if key == "" {
		return ""
	} else {
		value = os.Getenv(key)
		if len(value) == 0 {
			return ""
		}
		return value
	}
}

func GetEnvCors(key string) (fallback []string) {
	value := Getenv(key)
	return strings.Split(value, ",")
}
