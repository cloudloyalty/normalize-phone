package normalize_phone

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalizePhone(t *testing.T) {
	testCases := []struct {
		phoneNumber    string
		countryCode    string
		expectedResult string
		expectedErr    error
	}{
		// international
		{
			"+7 111 222 33 44",
			"ru",
			"+71112223344",
			nil,
		},
		{
			"+375 11 222 33 44",
			"ru",
			"+375112223344",
			nil,
		},
		{
			"+49 111 222 33 44",
			"ru",
			"+491112223344",
			nil,
		},
		{
			"+49 1111 222 33 44",
			"ru",
			"+4911112223344",
			nil,
		},
		{
			"+380 11 222 33 44",
			"ru",
			"+380112223344",
			nil,
		},
		{
			"+49 1111 222 33 44 5",
			"ru",
			"",
			ErrIncorrectPhone,
		},

		// ru
		{
			"8 111 222 33 44",
			"ru",
			"+71112223344",
			nil,
		},
		{
			"111 222 33 44",
			"ru",
			"+71112223344",
			nil,
		},
		{
			"9527162525",
			"ru",
			"+79527162525",
			nil,
		},
		{
			"79527162525",
			"ru",
			"+79527162525",
			nil,
		},

		// by
		{
			"80 11 222 33 44",
			"by",
			"+375112223344",
			nil,
		},
		{
			"11 222 33 44",
			"by",
			"+375112223344",
			nil,
		},

		// de
		{
			"0 111 222 33 44",
			"de",
			"+491112223344",
			nil,
		},
		{
			"111 222 33 44",
			"de",
			"+491112223344",
			nil,
		},
		{
			"0 1111 222 33 44",
			"de",
			"+4911112223344",
			nil,
		},
		{
			"1111 222 33 44",
			"de",
			"+4911112223344",
			nil,
		},

		// ua
		{
			"0 11 222 33 44",
			"ua",
			"+380112223344",
			nil,
		},
		{
			"11 222 33 44",
			"ua",
			"+380112223344",
			nil,
		},

		{
			"",
			"any country",
			"",
			nil,
		},
	}

	// Iterating over the test cases, call the function under test and assert the output.
	for i, testCase := range testCases {
		actualResult, err := NormalizePhone(testCase.countryCode, testCase.phoneNumber)
		assert.Equal(t, testCase.expectedResult, actualResult, fmt.Sprintf("case %d: %v", i, testCase))
		assert.Equal(t, testCase.expectedErr, err, fmt.Sprintf("case %d: %v", i, testCase))
	}
}
