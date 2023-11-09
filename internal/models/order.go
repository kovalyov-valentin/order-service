package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/bxcodec/faker"
)

type Order struct {
	OrderUID          string    `json:"order_uid"`
	TrackNumber       string    `json:"track_number"`
	Entry             string    `json:"entry"`
	Delivery          Delivery  `json:"delivery"`
	Payment           Payment   `json:"payment"`
	Items             []Item    `json:"items"`
	Locale            string    `json:"locale"`
	InternalSignature string    `json:"internal_signature"`
	CustomerID        string    `json:"customer_id"`
	DeliveryService   string    `json:"delivery_service"`
	ShardKey          string    `json:"shardkey"`
	SmID              int       `json:"sm_id"`
	DateCreated       time.Time `json:"date_created"`
	OofShard          string    `json:"oof_shard"`
}

func (o Order) Value() (driver.Value, error) {
	return json.Marshal(o)
}

func (o *Order) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}

func GetRandomOrderEncodedJson() json.RawMessage {
	ord := &Order{}
	err := faker.FakeData(ord)
	if err != nil {
		log.Fatal(err)
	}
	res, err := json.Marshal(ord)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
