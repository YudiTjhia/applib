package util

import (
	"strings"
)

func ContainsLower(pData string, pSearch string) bool {
	data := strings.ToLower(pData)
	search := strings.ToLower(pSearch)
	if strings.Contains(data, search) {
		return true
	}
	return false
}
