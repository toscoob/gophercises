package normalize

import (
	"errors"
	"fmt"
	"regexp"
)

const requiredLength = 10

func Normalize(in string) (string, error) {
	reg := regexp.MustCompile("\\D")

	processedString := reg.ReplaceAllString(in, "")

	if len(processedString) != requiredLength {
		return "", errors.New(fmt.Sprintf("number bad length: %s -> %s", in, processedString))
	}

	return processedString, nil
}
