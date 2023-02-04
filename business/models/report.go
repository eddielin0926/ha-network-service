package models

import "time"

type Report struct {
	Location string    `json:"location"`
	Date     time.Time `json:"date"`
	Count    int       `json:"count"`
	Material float32   `json:"material"`
	A        float32   `json:"a"`
	B        float32   `json:"b"`
	C        float32   `json:"c"`
	D        float32   `json:"d"`
}
