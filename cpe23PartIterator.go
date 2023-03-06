package cpe

import (
	"strings"

	"github.com/pkg/errors"
)

type Iterator[T any] interface {
	HasNext() bool
	Next() string
}

type Cpe23PartIterator struct {
	pos       int
	cpeString string
}

func (i *Cpe23PartIterator) HasNext() bool {
	return i.pos < len(i.cpeString)
}

func (i *Cpe23PartIterator) Cpe23PartIterator(cpeString string) error {
	if cpeString == "" || !strings.HasPrefix(cpeString, "cpe:2.3:") {
		return errors.New("Invalid 2.3 CPE value: " + i.cpeString)
	}
	i.cpeString = cpeString
	i.pos = 8

	return nil
}

func (i *Cpe23PartIterator) Next() string {
	end := 0
	for end = i.pos; end < len(i.cpeString); end++ {
		if i.cpeString[end:end+1] == ":" {
			break
		}
		if i.cpeString[end:end+1] == "\\" && end+1 < len(i.cpeString) {
			end += 1
		}
	}
	part := i.cpeString[i.pos:end]
	i.pos = end + 1

	return part
}
