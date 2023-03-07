package cpe

type Part string

const (
	APPLICATION      Part = "a"
	OPERATING_SYSTEM Part = "o"
	HARDWARE_DEVICE  Part = "h"
	ANY              Part = "*"
	NA               Part = "-"
)

func (p Part) String() string {
	switch p {
	case APPLICATION:
		return "a"
	case OPERATING_SYSTEM:
		return "o"
	case HARDWARE_DEVICE:
		return "h"
	case ANY:
		return "*"
	case NA:
		return "-"
	default:
		return ""
	}
}

func IsPart(part interface{}) bool {
	_, ok := part.(Part)
	return ok
}
