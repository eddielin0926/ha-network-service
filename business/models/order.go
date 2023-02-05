package models

type Order struct {
	Location  string `json:"location"`
	Timestamp string `json:"timestamp"`
	Data      Data   `json:"data"`
}
