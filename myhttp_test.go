package main

import (
	"testing"
)

func TestMakeRequestAndGetBodyMD5Checksum(t *testing.T) {
	testtable := []struct {
		url string // input
		formatted_url string
		md5str string
	}{
		{"www.reddit.com/r/aww.json", "http://www.reddit.com/r/aww.json", "5bceb0f13560b448376e5e3efd2ea7c3"},
		{"abc", "http://abc", "ERROR"},
		{"123456.123", "http://123456.123", "ERROR"},
	}

	for _, testcase := range testtable {
		url, checksum := MakeRequestAndGetBodyMD5Checksum(testcase.url)
		if url != testcase.formatted_url {
			t.Errorf("Formatted url was wrong, got: %s, expect: %s", url, testcase.formatted_url)
		}
		if checksum != testcase.md5str {
			t.Errorf("MD5 checksume wrong, got: %s, expect: %s", checksum, testcase.md5str)
		}
	}
}