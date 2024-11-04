package tools

import (
	"errors"
	"regexp"
)

func SplitFileSuffix(text string) (string, error) {
	re := regexp.MustCompile(`container:\s+(\w+)`)
	matchs := re.FindStringSubmatch(text)
	if matchs != nil {
		return matchs[1], nil
	} else {
		return "", errors.New("not match")
	}
}