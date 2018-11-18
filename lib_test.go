package main

import (
	"net/url"
	"testing"
)

func TestArrayFormat(t *testing.T) {
	s := make([]*url.URL, 0)
	if len(s) != 1 {
		// "s" below is a mistake; I meant to type "len(s)"
		//
		// I expect to get this error:
		// ./lib_test.go:15: Errorf format %d has arg s of wrong type []*net/url.URL
		t.Errorf("wrong length: want %d got %d", 1, s)
	}
}

func TestStringFormat(t *testing.T) {
	s := make([]string, 0)
	if len(s) != 1 {
		// "s" below is a mistake; I meant to type "len(s)"
		//
		// I expect to get this error:
		// ./lib_test.go:26: Errorf format %d has arg s of wrong type []string
		t.Errorf("wrong length: want %d got %d", 1, s)
	}
}
