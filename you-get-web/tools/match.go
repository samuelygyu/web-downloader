package tools

import (
	"errors"
	"regexp"
)

func SplitFileSuffix(text string) (string, string, error) {
	re1 := regexp.MustCompile(`title:\s+([\u4e00-\u9fa5\w]+)`)
	re2 := regexp.MustCompile(`container:\s+(\w+)`)
	matchs1 := re1.FindStringSubmatch(text)
	matchs2 := re2.FindStringSubmatch(text)
	if matchs2 != nil {
		if matchs1!= nil {
			return matchs2[1], matchs1[1], nil
		} else {
			return matchs2[1], "", nil
		}
	} else {
		return "", "", errors.New("not match")
	}
}