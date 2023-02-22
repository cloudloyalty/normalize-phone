package normalize_phone

import (
	"errors"
	"fmt"
	"strings"
)

var ErrIncorrectPhone = errors.New("incorrect phone")

// NormalizePhone changes phone string to international format without spaces or other symbols
func NormalizePhone(countryCode, phoneNumber string) (string, error) {
	if phoneNumber == "" {
		return phoneNumber, nil
	}

	// remove all non-digits
	phoneNumber = strings.Map(
		func(c rune) rune {
			if (c >= '0' && c <= '9') || c == '+' {
				return c
			}
			return -1
		},
		phoneNumber,
	)

	// international number
	if len(phoneNumber) > 0 && phoneNumber[0] == '+' {
		phoneNumber = phoneNumber[1:]
		for _, f := range countryPhoneFormat {
			if strings.HasPrefix(phoneNumber, f.Prefix) {
				for _, l := range f.Lengths {
					if len(phoneNumber) == l+len(f.Prefix) {
						return "+" + phoneNumber, nil
					}
				}
			}
		}
		return "", ErrIncorrectPhone
	}

	// country-specific checks
	f, ok := countryPhoneFormat[countryCode]
	if !ok {
		return "", fmt.Errorf("unknown country code '%s'", countryCode)
	}

	prefixMatches := strings.HasPrefix(phoneNumber, f.Prefix)
	for _, l := range f.Lengths {
		// with prefix, but without plus
		if prefixMatches && len(phoneNumber) == l+len(f.Prefix) {
			return "+" + phoneNumber, nil
		}
		// without prefix
		if len(phoneNumber) == l {
			return "+" + f.Prefix + phoneNumber, nil
		}
		// with trunk prefix
		for _, t := range f.TrunkPrefixes {
			if strings.HasPrefix(phoneNumber, t) && len(phoneNumber) == l+len(t) {
				return "+" + f.Prefix + phoneNumber[len(t):], nil
			}
		}
	}

	return "", ErrIncorrectPhone
}
