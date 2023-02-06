package server

import (
	"context"
	"encoding/base64"
	"strconv"

	"github.com/eddielin0926/ha-network-service/grpcpb/inventory"
)

type inventoryServer struct {
	inventory.UnimplementedInventoryServer
}

func NewInventoryServer() *inventoryServer {
	return &inventoryServer{}
}

func (s *inventoryServer) GetMaterial(ctx context.Context, in *inventory.Order) (*inventory.Material, error) {
	a, b, c, d := in.GetA(), in.GetB(), in.GetC(), in.GetD()
	material := a*3 + b*2 + c*4 + d*10
	signature := s.sign(a, b, c, d)
	return &inventory.Material{Material: material, Signature: signature}, nil
}

func (s *inventoryServer) sign(a, b, c, d int32) string {
	data := a + b + c + d
	strData := strconv.FormatInt(int64(data), 10)
	signature := base64.StdEncoding.EncodeToString([]byte(strData))
	return signature
}
