package cpe

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

func cpeUriToWellFormed(value string) string {
	switch {
	case len(value) == 0:
		return ""
	case value[0] == '/':
		ret := value[1:]
		if !IsPart(ret) {
			return ""
		}
		return ret
	case value[len(value)-1] == '/':
		return value[:len(value)-1]
	default:
		return value
	}
}

// TODO 判断是否用得上
func cpeUriToWellFormed222(value string) (string, error) {
	if value == "" {
		return "*", nil
	} else if value == "*" {
		return "*", nil
	}

	bytes := []byte(strings.ToLower(value))
	sb := strings.Builder{}

	for x := 0; x < len(bytes); x++ {
		r, size := utf8.DecodeRune(bytes[x:])
		if r == utf8.RuneError && size == 1 {
			return "", errors.New("invalid utf8 sequence")
		}

		if (r >= '0' && r <= '9') || (r >= 'a' && r <= 'z') {
			sb.WriteRune(r)
		} else if r == '_' || r == '.' || r == '-' {
			sb.WriteRune(r)
		} else if r == '%' {
			if (2 + x) >= len(bytes) {
				return "", errors.New("invalid CPE URI component - ends with a single percent")
			}
			decoded, err := strconv.ParseUint(string(bytes[x+1:x+3]), 16, 16)
			if err != nil {
				return "", err
			}
			switch decoded {
			case 1:
				sb.WriteRune('?')
			case 2:
				sb.WriteRune('*')
			default:
				sb.WriteRune('\\')
				sb.WriteRune(rune(decoded))
			}
		} else {
			sb.WriteRune('\\')
			sb.WriteRune(r)
		}

	}

	return sb.String(), nil
}
