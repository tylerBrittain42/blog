package validator

import (
	"errors"
	"os"
	"path/filepath"
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

func IsAccessible(dir string, fileName string) (bool, error) {

	filePath := filepath.Join(dir, fileName)

	_, err := os.Stat(filePath)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
