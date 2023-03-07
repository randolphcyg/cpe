package cpe

var Part = map[string]struct{}{
	"a": {},
	"o": {},
	"h": {},
}

func IsPart(k string) bool {
	_, ok := Part[k]
	return ok
}
