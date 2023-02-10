package server

import (
	"context"
	"sync"

	"github.com/eddielin0926/ha-network-service/grpcpb/storage"
)

type storageServer struct {
	storage.UnimplementedStorageServer
	records []*storage.Record
	mu      *sync.Mutex
}

func NewStorageServer() *storageServer {
	return &storageServer{
		mu: &sync.Mutex{},
	}
}

func (s *storageServer) SaveRecord(ctx context.Context, in *storage.Record) (*storage.Response, error) {
	r := *in
	s.mu.Lock()
	s.records = append(s.records, &r)
	s.mu.Unlock()
	return &storage.Response{Status: storage.Status_SUCCESS}, nil
}

func (s *storageServer) GetRecords(ctx context.Context, in *storage.Query) (*storage.RecordsResponse, error) {
	result := s.filterRecords(in.GetLocation(), in.GetDate())
	return &storage.RecordsResponse{Records: result}, nil
}

func (s *storageServer) GetReport(ctx context.Context, in *storage.Query) (*storage.Report, error) {
	report := &storage.Report{}
	data := s.filterRecords(in.GetLocation(), in.GetDate())
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

func (s *storageServer) filterRecords(location string, date string) []*storage.Record {
	var result []*storage.Record
	for _, e := range s.records {
		if e.Location == location && s.isEqualDate(e.Timestamp, date) {
			result = append(result, e)
		}
	}
	return result
}

func (s *storageServer) isEqualDate(date1, date2 string) bool {
	return date1[:10] == date2[:10]
}
