package cpe

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
