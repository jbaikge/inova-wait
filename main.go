package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const WaitURL = "https://www.inovaspine.org/emergency-room-wait-times/"

func main() {
	var filename string
	var ids []int

	// Path to store history
	if filename = os.Getenv("STORE_FILE"); filename == "" {
		log.Fatalf("STORE_FILE not defined")
	}

	// Convert space-separated IDs to int array
	if idEnv := os.Getenv("TRACK_IDS"); idEnv == "" {
		log.Fatalf("TRACK_IDS not defined, nothing to track")
	} else {
		ids = make([]int, 0, 16)
		for _, id := range strings.Fields(idEnv) {
			n, err := strconv.Atoi(id)
			if err != nil {
				log.Fatalf("trouble converting %v to int", id)
			}
			ids = append(ids, n)
		}
	}

	html, err := GetHTML(WaitURL)
	if err != nil {
		log.Fatalf("Error fetching HTML: %v", err)
	}

	waits, err := ParseWaits(html)
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	last, err := GetLatestWaits(filename)
	if err != nil {
		log.Fatalf("Error fetching latest waits: %v", err)
	}

	if err := StoreWaits(filename, waits); err != nil {
		log.Fatalf("Error storing data: %v", err)
	}

	if len(last) == 0 {
		log.Printf("No previous data, exiting")
		return
	}

	// for _, wait := range waits {
	// 	fmt.Printf("%2d  %2dmin  %s\n", wait.Id, wait.Minutes, wait.Location)
	// }

	// body := "Inova Fairfax Hospital decreased to 14min\nInova Emergency Room - Fairfax increased to 7min"
	// if err := SendMessage(body); err != nil {
	// 	log.Fatalf("Error sending message: %v", err)
	// }
}
