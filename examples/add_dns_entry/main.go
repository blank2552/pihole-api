package main

import (
	"log"
	"os"

	piholeapi "github.com/blank2552/pihole-api"
)

func main() {
	baseURL := os.Getenv("PIHOLE_BASE_URL")
	password := os.Getenv("PIHOLE_PASSWORD")
	address := os.Getenv("PIHOLE_DNS_ADDRESS")
	hostname := os.Getenv("PIHOLE_DNS_HOSTNAME")
	if baseURL == "" || password == "" || address == "" || hostname == "" {
		log.Fatal("set PIHOLE_BASE_URL, PIHOLE_PASSWORD, PIHOLE_DNS_ADDRESS, and PIHOLE_DNS_HOSTNAME")
	}

	client := piholeapi.NewClient(baseURL, true)
	if err := client.Login(password); err != nil {
		log.Fatal(err)
	}

	if err := client.AddDNSEntry(address, hostname); err != nil {
		log.Fatal(err)
	}
}
