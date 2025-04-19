package graph

import (
	"encoding/json"
	"fmt"
	"log"
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

type DbConn interface {
	Connect() *DbStore
	Orders() []Record
}

type DbStore struct {
	records []Record
}

var (
	dataFile = "data.json"
	fullPath = fmt.Sprintf("/opt/data/%v", dataFile)
)

func (db *DbStore) Connect() *DbStore {
	log.Println("Connected")
	if err := db.Load(fullPath); err != nil {
		if err := db.Load("data.json"); err != nil {
			panic(err)
		}
	}
	return db
}

func (db *DbStore) Load(dataPath string) error {
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return err
	}
	var records []Record
	if err := json.Unmarshal(data, &records); err != nil {
		return err
	}
	db.records = records
	return nil
}

func (db *DbStore) Orders() []Record {
	return db.records
}
