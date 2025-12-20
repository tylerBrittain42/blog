package helper

import "regexp"

func IsAlphaNumeric(input string) (bool, error) {
	// regex: ^[a-zA-Z0-9]*$
	regexString := "^[a-zA-Z0-9]*$"
	isMatch, err := regexp.MatchString(regexString, input)
	if err != nil {
		return false, err
	}

	return isMatch, nil
}

func IsAccessible(input string) (bool, error) {
	return false, nil
}
