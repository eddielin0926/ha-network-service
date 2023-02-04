package models

import "time"

type Record struct {
	Location  string    `json:"location"`
	Timestamp time.Time `json:"timestamp"`
	Signature string    `json:"signature"`
	Material  float32   `json:"material"`
	A         float32   `json:"a"`
	B         float32   `json:"b"`
	C         float32   `json:"c"`
	D         float32   `json:"d"`
}
