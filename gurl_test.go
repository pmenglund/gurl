package main

import "testing"

func TestParseURL(t *testing.T) {
	u := parseURL("ftp://foo.bar")
	if u.Scheme != "ftp" {
		t.Error("oops")
	}
}
