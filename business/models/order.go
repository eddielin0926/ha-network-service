package models

import "time"

type Order struct {
	Location  string    `json:"location"`
	Timestamp time.Time `json:"timestamp"`
	Data      Data      `json:"data"`
}
