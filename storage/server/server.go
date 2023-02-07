package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/eddielin0926/ha-network-service/grpcpb/storage"
	"github.com/eddielin0926/ha-network-service/storage/models"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/gorm"
)

type storageServer struct {
	storage.UnimplementedStorageServer
	DB *gorm.DB
}

func NewStorageServer(db *gorm.DB) *storageServer {
	return &storageServer{DB: db}
}

func (s *storageServer) SaveRecord(ctx context.Context, in *storage.Record) (*storage.Response, error) {
	s.DB.Create(&models.Record{
		Location:  in.Location,
		Timestamp: in.Timestamp,
		Signature: in.Signature,
		Material:  in.Material,
		A:         in.A,
		B:         in.B,
		C:         in.C,
		D:         in.D,
	})
	return &storage.Response{Status: storage.Status_SUCCESS}, nil
}

func (s *storageServer) GetRecords(ctx context.Context, in *storage.Query) (*storage.RecordsResponse, error) {
	var data []models.Record
	s.DB.Where("location = ? AND timestamp LIKE ?", in.GetLocation(), in.GetDate()+"%").Find(&data)
	fmt.Println(data)
	u := protojson.UnmarshalOptions{
		AllowPartial:   true,
		DiscardUnknown: true,
	}

	var records []*storage.Record
	for _, e := range data {
		r := &storage.Record{}
		bin, _ := json.Marshal(e)
		u.Unmarshal(bin, r)
		records = append(records, r)
	}

	return &storage.RecordsResponse{Records: records}, nil
}

func (s *storageServer) GetReport(ctx context.Context, in *storage.Query) (*storage.Report, error) {
	report := &storage.Report{}
	var data []models.Record
	s.DB.Where("location = ? AND timestamp LIKE ?", in.GetLocation(), in.GetDate()+"%").Find(&data)
	for _, e := range data {
		report.A += e.A
		report.B += e.B
		report.C += e.C
		report.D += e.D
		report.Material += e.Material
	}
	report.Location = in.GetLocation()
	report.Date = in.GetDate()
	report.Count = int32(len(data))
	return report, nil
}
