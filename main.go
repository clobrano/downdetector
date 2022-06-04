package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

// Connect attemps to connect to a url. It returns with error if it does not succeed in 5 seconds
func Connect(url string, t time.Duration) error {
	client := http.Client{
		Timeout: t,
	}
	_, err := client.Get(url)
	if err != nil {
		log.Printf("could not connect to %s: %v\n", url, err)
		return err
	}
	log.Printf("successfully connect to %s\n", url)
	return err
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var config_file_name string
	var logging_file_name string

	flag.StringVar(&config_file_name, "configure", "configure.yml", "The configuration file name")
	flag.StringVar(&logging_file_name, "logfile", "logger.log", "The logging file name")

	flag.Parse()

	f, err := os.OpenFile(logging_file_name, os.O_CREATE|os.O_RDWR, 0666)
	defer f.Close()
	if err != nil {
		log.Panicf("could not open logging file: %+v\n", err)
	}
	log.SetOutput(f)

	for {
		config, err := NewConfiguration(config_file_name)
		if err != nil {
			log.Fatal(err)
		}
		for _, url := range config.Urls {
			err = Connect(url, config.Timeout.Value)
			if err != nil {
				msg := fmt.Sprintf("could not reach %s: error %v",
					url, err)

				if config.Action.On_failure.Telegram {
					notify, err := NewTelegramMessage(msg)
					if err != nil {
						log.Printf("could not create Telegram notifier: %v", err)
					} else {
						err = notify.Send()
						if err != nil {
							log.Fatalf("could not notify via Telegram: %v", err)
						}
					}
				}
			}
		}
		delay := time.Duration(config.Schedule.Hour)*time.Hour +
			time.Duration(config.Schedule.Minute)*time.Minute +
			time.Duration(config.Schedule.Second)*time.Second
		time.Sleep(delay)
	}
}
