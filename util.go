package cpe

import (
	"strings"
	"unicode/utf8"
)

func RegionMatches(srcStr string, ignoreCase bool, toffset int, other string, ooffset int, len int) bool {
	if ignoreCase {
		srcStr = strings.ToLower(srcStr)
		other = strings.ToLower(other)
	}

	var strTmp string
	if toffset+len <= utf8.RuneCountInString(srcStr) {
		strTmp = srcStr[toffset : toffset+len]
	}

	if other[ooffset:] == strTmp {
		return true
	}

	return false
}
