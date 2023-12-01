package utils

import (
	"strings"
)

func CleanerTxt(txt string) string {
	t := strings.TrimSpace(txt)
	fields := strings.Fields(t)
	return strings.Join(fields, " ")
}
