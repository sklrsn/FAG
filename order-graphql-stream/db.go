package main

import (
	"encoding/json"
	"os"
	"time"
)

type Record struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Quantity int       `json:"quantity"`
	Created  time.Time `json:"created"`
	User     struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	} `json:"user"`
	Payment struct {
		ID      string    `json:"id"`
		Amount  int       `json:"amount"`
		Created time.Time `json:"created"`
	} `json:"payment"`
	Shipping struct {
		ID      string    `json:"id"`
		Created time.Time `json:"created"`
		Address string    `json:"address"`
	} `json:"shipping"`
}

func PrepareDB() []Record {
	data, err := os.ReadFile("/opr/data/data.json")
	if err != nil {
		panic(err)
	}

	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		panic(err)
	}

	return records
}
