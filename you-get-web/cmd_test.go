package main

import (
	"testing"
)

func TestGenerateCommandProxy(t *testing.T) {
	cmd := generateCommand("www.baidu.com", "127.0.0.1:1111", "")
	want := "you-get www.baidu.com -x 127.0.0.1:1111 -o /Temp/you-get/download"
	if cmd != want {
		t.Errorf("generateCommand() = %v, want %v", cmd, want)
	}
}
func TestGenerateCommandCookies(t *testing.T) {
	cmd := generateCommand("www.baidu.com", "", "test.cookies.txt")
	want := "you-get www.baidu.com -c test.cookies.txt -o /Temp/you-get/download"
	if cmd != want {
		t.Errorf("generateCommand() = %v, want %v", cmd, want)
	}
}
func TestGenerateCommandProxyAndCookies(t *testing.T) {
	cmd := generateCommand("www.baidu.com", "127.0.0.1:1111", "test.cookies.txt")
	want := "you-get www.baidu.com -x 127.0.0.1:1111 -c test.cookies.txt -o /Temp/you-get/download"
	if cmd != want {
		t.Errorf("generateCommand() = %v, want %v", cmd, want)
	}
}
