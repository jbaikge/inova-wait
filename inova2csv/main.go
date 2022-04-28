package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type dataStore struct {
	Locations []string
	Minutes   [][]int64
	Times     []time.Time
}

func main() {
	var store dataStore
	if err := json.NewDecoder(os.Stdin).Decode(&store); err != nil {
		log.Fatalf("Decode error: %v", err)
	}

	w := csv.NewWriter(os.Stdout)

	headers := append([]string{"Time"}, store.Locations...)
	if err := w.Write(headers); err != nil {
		log.Fatalf("Header write error: %v", err)
	}

	for i, mins := range store.Minutes {
		t := store.Times[i]
		record := []string{t.Local().Format("1/2/2006 3:04pm")}
		for _, min := range mins {
			record = append(record, fmt.Sprint(min))
		}
		if err := w.Write(record); err != nil {
			log.Fatalf("Record write error: %v", err)
		}
	}

	w.Flush()
}
