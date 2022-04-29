package main

import (
	"encoding/json"
	"errors"
	"log"
	"os"
	"time"
)

type dataStore struct {
	Locations []string
	Minutes   [][]int
	Times     []time.Time
}

func getData(filename string) (store dataStore, err error) {
	f, err := os.Open(filename)
	if errors.Is(err, os.ErrNotExist) {
		return store, nil
	}

	if err != nil {
		return
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&store)
	return
}

func GetLatestWaits(filename string) (waits []Wait, err error) {
	store, err := getData(filename)
	if err != nil {
		return
	}

	waits = make([]Wait, len(store.Locations))
	last := len(store.Minutes) - 1
	if len(store.Minutes[last]) != len(store.Locations) {
		log.Printf("Error: Last minutes only has %d items! Using previous record.", len(store.Minutes[last]))
		last--
	}
	for i := range store.Locations {
		waits[i].Id = i + 1
		waits[i].Location = store.Locations[i]
		waits[i].Minutes = store.Minutes[last][i]
	}
	return
}

func StoreWaits(filename string, waits []Wait) (err error) {
	store, err := getData(filename)
	if err != nil {
		return
	}

	if len(store.Locations) == 0 {
		store.Locations = make([]string, len(waits))
		for i := range waits {
			store.Locations[i] = waits[i].Location
		}

		store.Minutes = make([][]int, 0, 1)
		store.Times = make([]time.Time, 0, 1)
	}
	
	minutes := make([]int, len(waits))
	for i := range waits {
		minutes[i] = waits[i].Minutes
	}
	store.Minutes = append(store.Minutes, minutes)

	store.Times = append(store.Times, time.Now())

	f, err := os.Create(filename)
	if err != nil {
		return
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(&store)
}
