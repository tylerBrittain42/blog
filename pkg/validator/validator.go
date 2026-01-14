package validator

import (
	"errors"
	"os"
	"regexp"
)

func IsAlphaNumeric(input string) (bool, error) {
	// regex: ^[a-zA-Z0-9]*$
	regexString := "^[a-zA-Z0-9]*$"
	isMatch, err := regexp.MatchString(regexString, input)
	if err != nil {
		return false, err
	}

	return isMatch, nil
}

func IsAccessible(filePath string) (bool, error) {

	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
