package utils

import (
	"fmt"
	"regexp"

	"github.com/joomcode/errorx"
)

func CreateValidateEntityNames(
	inclusions []string,
	exclusions []string,
) (ValidateEntityName, error) {
	inclusionRegexes, err := compileRegexes(inclusions)

	if err != nil {
		return nil, err
	}

	validateInclusion := func(entityName string) bool {
		if len(inclusionRegexes) == 0 {
			return true
		}

		for _, regex := range inclusionRegexes {
			if regex.MatchString(entityName) {
				return true
			}
		}

		return false
	}

	exclusionRegexes, err := compileRegexes(exclusions)

	if err != nil {
		return nil, err
	}

	validateExclusion := func(entityName string) bool {
		if len(exclusionRegexes) == 0 {
			return true
		}

		for _, regex := range exclusionRegexes {
			if regex.MatchString(entityName) {
				return false
			}
		}

		return true
	}

	validateEntityName := func(entityName string) bool {
		return validateInclusion(entityName) && validateExclusion(entityName)
	}

	return validateEntityName, nil
}

func compileRegexes(patterns []string) ([]*regexp.Regexp, error) {
	regexes := make([]*regexp.Regexp, 0)

	for _, pattern := range patterns {
		regex, err := regexp.Compile(pattern)

		if err != nil {
			msg := fmt.Sprintf("failed to compile regex '%s'", pattern)

			return nil, errorx.IllegalFormat.Wrap(err, msg)
		}

		regexes = append(regexes, regex)
	}

	return regexes, nil
}

type ValidateEntityName func(entityName string) bool
