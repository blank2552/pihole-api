package main

import (
	"log"
	"os"

	piholeapi "github.com/blank2552/pihole-api"
)

func main() {
	baseURL := os.Getenv("PIHOLE_BASE_URL")
	password := os.Getenv("PIHOLE_PASSWORD")
	if baseURL == "" || password == "" {
		log.Fatal("set PIHOLE_BASE_URL and PIHOLE_PASSWORD")
	}

	client := piholeapi.NewClient(baseURL, true)
	if err := client.Login(password); err != nil {
		log.Fatal(err)
	}

	if err := client.UpdateGravity(); err != nil {
		log.Fatal(err)
	}
}
