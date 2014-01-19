package main

import (
	"strings"
)

func lowercase(str string) string {
	return strings.ToLower(str[:1]) + str[1:]
}

func capitalize(str string) string {
	return strings.ToUpper(str[:1]) + str[1:]
}
