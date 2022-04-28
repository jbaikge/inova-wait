package main

import (
	"fmt"
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

	currentWaits, err := ParseWaits(html)
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	lastWaits, err := GetLatestWaits(filename)
	if err != nil {
		log.Fatalf("Error fetching latest waits: %v", err)
	}

	if err := StoreWaits(filename, currentWaits); err != nil {
		log.Fatalf("Error storing data: %v", err)
	}

	if len(lastWaits) == 0 {
		log.Printf("No previous data, exiting")
		return
	}

	// Determine wait time changes
	haveChanges := false
	messageLines := make([]string, len(ids))
	for i, id := range ids {
		var last, current Wait
		for _, wait := range lastWaits {
			if wait.Id == id {
				last = wait
				break
			}
		}
		for _, wait := range currentWaits {
			if wait.Id == id {
				current = wait
				break
			}
		}

		dir := "remains at"
		if last.Minutes > current.Minutes {
			dir = "decreased to"
			haveChanges = true
		} else if last.Minutes < current.Minutes {
			dir = "increased to"
			haveChanges = true
		}

		messageLines[i] = fmt.Sprintf("%s %s %dmin", current.Location, dir, current.Minutes)
	}

	if !haveChanges {
		log.Printf("No time changes since last check. Exiting")
		return
	}

	body := strings.Join(messageLines, "\n\n")
	if err := SendMessage(body); err != nil {
		log.Fatalf("Error sending message: %v", err)
	}
}
