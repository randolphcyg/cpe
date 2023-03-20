package cpe

import (
	"strings"

	"github.com/pkg/errors"
)

var (
	ErrCPENonstandard               = errors.New("the CPE string specified does not conform to the CPE 2.2 or 2.3 specification")
	ErrCPEStrEmptyOrNull            = errors.New("CPE String is null ir empty")
	ErrInvalidPart                  = errors.New("CPE String contains a invalid part!")
	ErrInvalidPartTooManyComponents = errors.New("CPE String contains a invalid part!too many components.")
	ErrInvalidPartTooFewComponents  = errors.New("CPE String contains a invalid part!too few components.")
	ErrInvalidCPE                   = errors.New("Invalid CPE")
)

// CPE represents a Common Platform Enumeration
type CPE struct {
	Part     string
	Vendor   string
	Product  string
	Version  string
	Update   string
	Edition  string
	Language string
	// New field in CPE 2.3
	SwEdition string
	TargetSw  string
	TargetHw  string
	Other     string
}

func NewCPE() *CPE {
	return &CPE{
		Part:      "",
		Vendor:    "",
		Product:   "",
		Version:   "",
		Update:    "",
		Edition:   "",
		Language:  "",
		SwEdition: "",
		TargetSw:  "",
		TargetHw:  "",
		Other:     "",
	}
}

// ParseCPE
/**
 * Parses a CPE String into an object with the option of parsing CPE 2.2 URI
 * strings in lenient mode - allowing for CPE values that do not adhere to
 * the specification.
 *
 * @param cpeString the CPE string to parse
 * @return the CPE object represented by the given cpeString
 * @throws error thrown if the cpeString is invalid
 */
func ParseCPE(cpeString string) (cpe *CPE, err error) {
	if len(cpeString) == 0 {
		return cpe, ErrCPEStrEmptyOrNull
	} else if RegionMatches(cpeString, false, 0, "cpe:/", 0, 5) {
		return parseCPE22(cpeString)
	} else if RegionMatches(cpeString, false, 0, "cpe:2.3:", 0, 8) {
		return parseCPE23(cpeString)
	} else {
		return cpe, ErrCPENonstandard
	}
}

/**
 * Parses a CPE 2.2 URI.
 *
 * @param cpeString the CPE string to parse
 * @return the CPE object represented by the cpeString
 * @throws error thrown if the cpeString is invalid
 */
func parseCPE22(cpeString string) (cpe *CPE, err error) {
	if len(cpeString) == 0 {
		return cpe, ErrCPEStrEmptyOrNull
	}
	parts := strings.Split(cpeString, ":")
	if len(parts) <= 1 || len(parts) > 8 {
		return cpe, ErrInvalidCPE
	}
	if len(parts[1]) != 2 {
		return cpe, ErrInvalidPart
	}

	cpe = NewCPE()
	cpe.Part = cpeUriToWellFormed(parts[1])
	if len(cpe.Part) == 0 {
		return cpe, ErrInvalidPart
	}
	if len(parts) > 2 {
		cpe.Vendor = cpeUriToWellFormed(parts[2])
	}
	if len(parts) > 3 {
		cpe.Product = cpeUriToWellFormed(parts[3])
	}
	if len(parts) > 4 {
		cpe.Version = cpeUriToWellFormed(parts[4])
	}
	if len(parts) > 5 {
		cpe.Update = cpeUriToWellFormed(parts[5])
	}
	if len(parts) > 6 {
		edition := cpeUriToWellFormed(parts[6])
		err = cpe.unpackEdition(edition)
		if err != nil {
			return cpe, err
		}
	}
	if len(parts) > 7 {
		cpe.Language = cpeUriToWellFormed(parts[7])
	}

	return cpe, nil
}

/**
 * In a CPE 2.2 URI the new fields from CPE 2.3 may be "packed" into the
 * edition field. If present, each field will be preceeded by a '~'.
 * Example, "~edition~swEdition~targetSw~targetHw~other".
 *
 * @param edition the edition string to unpack
 * @param cpe a reference to the CPE Builder to unpack the edition into
 */
func (cpe *CPE) unpackEdition(edition string) error {
	if len(edition) == 0 {
		return errors.New("edition is null")
	}

	unpacked := strings.Split(edition, "~")
	if strings.HasPrefix(edition, "~") {
		if len(unpacked) > 1 {
			cpe.Edition = unpacked[1]
		}
		if len(unpacked) > 2 {
			cpe.SwEdition = unpacked[2]
		}
		if len(unpacked) > 3 {
			cpe.TargetSw = unpacked[3]
		}
		if len(unpacked) > 5 {
			cpe.TargetHw = unpacked[4]
		}
		if len(unpacked) > 5 {
			cpe.Other = unpacked[5]
		}
		if len(unpacked) > 6 {
			return errors.New("invalid packed edition")
		}
	} else {
		cpe.Edition = edition
	}

	return nil
}

func parseCPE23(cpeString string) (cpe *CPE, err error) {
	if len(cpeString) == 0 {
		return cpe, ErrCPEStrEmptyOrNull
	}

	iter := Cpe23PartIterator{}
	err = iter.Cpe23PartIterator(cpeString)
	if err != nil {
		return cpe, err
	}

	cpe = NewCPE()
	if iter.HasNext() {
		cpe.Part = iter.Next()
		if !IsPart(cpe.Part) {
			return cpe, ErrInvalidPart
		}
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.Vendor = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.Product = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.Version = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.Update = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.Edition = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.Language = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.SwEdition = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.TargetSw = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.TargetHw = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		cpe.Other = iter.Next()
	} else {
		return cpe, ErrInvalidPartTooFewComponents
	}

	if iter.HasNext() {
		return cpe, ErrInvalidPartTooManyComponents
	}

	return cpe, nil
}
