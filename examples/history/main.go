package main

import (
	"fmt"
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

	history, err := client.GetHistory()
	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range history.History {
		fmt.Printf(
			"%d total=%d cached=%d blocked=%d forwarded=%d\n",
			entry.Timestamp,
			entry.Total,
			entry.Cached,
			entry.Blocked,
			entry.Forwarded,
		)
	}
}
