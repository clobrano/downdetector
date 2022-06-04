package main

import (
	"testing"
	"time"
)

var data = `
urls:
  - websiteA.com
  - websiteB.com
timeout: 1
schedule:
  hour: 1
  minute: 2
  second: 3
action:
  on-success:
    telegram: true
    email: false
  on-failure:
    telegram: true
    email: false
`

func TestConfiguration(t *testing.T) {
	var config Configuration
	err := config.Parse([]byte(data))
	if err != nil {
		t.FailNow()
	}
	if config.Urls[0] != "websiteA.com" {
		t.Fatalf("urls 0 is '%s'", config.Urls[0])
	}
	if config.Urls[1] != "websiteB.com" {
		t.Fatalf("urls 1 is '%s'", config.Urls[1])
	}
	if config.Timeout.Value != 1*time.Second {
		t.Fatalf("timeout is %+v", config.Timeout)
	}
	if config.Schedule.Hour != 1 {
		t.Fatalf("schedule hour is %d", config.Schedule.Hour)
	}
	if config.Schedule.Minute != 2 {
		t.Fatalf("schedule minute is %d", config.Schedule.Minute)
	}
	if config.Schedule.Second != 3 {
		t.Fatalf("schedule second is %d", config.Schedule.Second)
	}
	if config.Action.On_success.Telegram != true {
		t.Fatalf("action on-success service Telegram is '%v'", config.Action.On_success.Telegram)
	}
	if config.Action.On_success.Email != false {
		t.Fatalf("action on-success service Email is '%v'", config.Action.On_success.Telegram)
	}
}
