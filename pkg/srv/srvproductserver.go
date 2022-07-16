package srv

import (
	"TreinoTest/pkg/pb"
	"context"
	"log"
)

type ProductServer struct {
	pb.UnimplementedProductsServiceServer
}

func (service *ProductServer) Create(ctx context.Context, req *pb.ProductRequest) (*pb.ProductResponse, error) {
	log.Println("Mensagem recebida", req.Name)
	response := &pb.ProductResponse{Name: "Eita"}
	return response, nil
}