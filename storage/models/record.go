package models

import "gorm.io/gorm"

type Record struct {
	gorm.Model
	Location  string  `json:"location"`
	Timestamp string  `json:"timestamp"`
	Signature string  `json:"signature"`
	Material  int32   `json:"material"`
	A         float32 `json:"a"`
	B         float32 `json:"b"`
	C         float32 `json:"c"`
	D         float32 `json:"d"`
}
