package models

type Record struct {
	Location  string  `json:"location"`
	Timestamp string  `json:"timestamp"`
	Signature string  `json:"signature"`
	Material  int32   `json:"material"`
	A         float32 `json:"a"`
	B         float32 `json:"b"`
	C         float32 `json:"c"`
	D         float32 `json:"d"`
}

type Records struct {
	Records []Record
}
