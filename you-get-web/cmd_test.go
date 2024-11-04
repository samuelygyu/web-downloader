package main

import (
	"fmt"
	"regexp"
	"testing"
)

func TestRegexp(t *testing.T) {
	text := "site:                Bilibili\ntitle:               徐静雨玩恐怖游戏心率飙升180，荣获B站软人奖\nstream:\n- format:       [7mdash-flv-AVC0m\ncontainer:     mp4\nquality:       高清 1080P avc1.640032"
	re := regexp.MustCompile(`container:\s+(\w+)`)
	matchs := re.FindStringSubmatch(text)

	if matchs != nil {
		fmt.Println(matchs[1])
	} else {
		t.Error("not match")
	}
}
