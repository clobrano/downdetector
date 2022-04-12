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

func TestConfigurationFromData(t *testing.T) {
	var config Configuration
	err := config.Parse([]byte(data))
	if err != nil {
		t.FailNow()
	}
	if config.Urls[0] != "websiteA.com" {
		t.Fatalf("urls 0 is '%s'", config.Urls[0])
		t.Fail()
	}
	if config.Urls[1] != "websiteB.com" {
		t.Fatalf("urls 1 is '%s'", config.Urls[1])
		t.Fail()
	}
	if config.Timeout.Value != 1*time.Second {
		t.Fatalf("timeout is %+v", config.Timeout)
		t.Fail()
	}
	if config.Schedule.Hour != 1 {
		t.Fatalf("schedule hour is %d", config.Schedule.Hour)
		t.Fail()
	}
	if config.Schedule.Minute != 2 {
		t.Fatalf("schedule minute is %d", config.Schedule.Minute)
		t.Fail()
	}
	if config.Schedule.Second != 3 {
		t.Fatalf("schedule second is %d", config.Schedule.Second)
		t.Fail()
	}
	if config.Action.On_success.Telegram != true {
		t.Fatalf("action on-success service Telegram is '%v'", config.Action.On_success.Telegram)
		t.Fail()
	}
	if config.Action.On_success.Email != false {
		t.Fatalf("action on-success service Email is '%v'", config.Action.On_success.Telegram)
		t.Fail()
	}
}

func TestConfigurationFromFileName(t *testing.T) {
	config, err := NewConfiguration("configure_example.yml")
	if err != nil {
		t.Fatalf("%v", err)
		t.Fail()
	}
	if (*config).Urls[0] != "www.google.com" {
		t.Fatalf("url 0 is not www.google.com")
		t.Fail()
	}
	if (*config).Urls[1] != "www.duckduckgo.com" {
		t.Fatalf("url 1 is not www.duckduckgo.com")
		t.Fail()
	}
	if (*config).Schedule.Hour != 1 {
		t.Fatalf("schedule Hour is %d", (*config).Schedule.Hour)
		t.Fail()
	}
	if (*config).Schedule.Minute != 2 {
		t.Fatalf("schedule Minute is %d", (*config).Schedule.Minute)
		t.Fail()
	}
	if (*config).Schedule.Second != 3 {
		t.Fatalf("schedule Second is %d", (*config).Schedule.Second)
		t.Fail()
	}
	if (*config).Action.On_success.Telegram != true {
		t.Fatalf("Action OnSuccess Telegram is %v", (*config).Action.On_success.Telegram)
		t.Fail()
	}
}
