package main

import (
	"io/ioutil"
	"log"
	"time"

	"gopkg.in/yaml.v2"
)

type Timeout struct {
	Value time.Duration
}

// UnmarshalYAML unmarshals an integer converting it from Nanoseconds to Seconds
func (t *Timeout) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var seconds int64
	err := unmarshal(&seconds)
	if err != nil {
		return err
	}
	t.Value = time.Duration(seconds * 1000 * 1000 * 1000)
	return nil
}

type notificationSystem struct {
	Telegram bool `yaml:"telegram"`
	Email    bool `yaml:"email"`
}

type action struct {
	On_success notificationSystem `yaml:"on-success"`
	On_failure notificationSystem `yaml:"on-failure"`
}

type schedule struct {
	Hour   int `yaml:"hour"`
	Minute int `yaml:"minute"`
	Second int `yaml:"second"`
}

type Configuration struct {
	Urls     []string `yaml:"urls"`
	Schedule schedule `yaml:"schedule"`
	Action   action   `yaml:"action"`
	Timeout  Timeout  `yaml:"timeout"`
}

// Parse reads a yaml configuration file
func (c *Configuration) Parse(data []byte) error {
	return yaml.Unmarshal(data, c)
}

// NewConfiguration creates a new Configuration from yaml file
func NewConfiguration(filename string) (*Configuration, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var config Configuration
	if err := config.Parse(data); err != nil {
		log.Fatalf("could not read configuration from file '%s': error %+v", filename, err)
	}
	return &config, err
}
