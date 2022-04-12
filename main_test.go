package main

import (
	"testing"
	"time"
)

func TestConnect(t *testing.T) {
	valid_url := "https://www.google.com"
	err := Connect(valid_url, time.Second)
	if err != nil {
		t.Fail()
	}

	invalid_url := "https://www.foo.com"
	err = Connect(invalid_url, time.Second)
	if err == nil {
		t.Fail()
	}
}
