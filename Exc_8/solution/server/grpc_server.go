package server

import (
	"context"
	"exc8/pb"
	"net"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type GRPCService struct {
	pb.UnimplementedOrderServiceServer
	drinks []*pb.Drink
	orders []*pb.Order
}

func StartGrpcServer() error {
	// Create a new gRPC server.
	srv := grpc.NewServer()
	// Create grpc service
	grpcService := &GRPCService{
		drinks: []*pb.Drink{
			{Id: 1, Name: "Spritzer", Price: 2.0, Description: "Wine with soda"},
			{Id: 2, Name: "Beer", Price: 3.0, Description: "Hagenberger Gold"},
			{Id: 3, Name: "Coffee", Price: 0.0, Description: "Mifare isn't that secure"},
		},
		orders: make([]*pb.Order, 0),
	}
	// Register our service implementation with the gRPC server.
	pb.RegisterOrderServiceServer(srv, grpcService)
	// Serve gRPC server on port 4000.
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		return err
	}
	err = srv.Serve(lis)
	if err != nil {
		return err
	}
	return nil
}

// todo implement functions
// ListDrinks implements: rpc ListDrinks(Empty) returns (ListDrinksResponse)
func (s *GRPCService) ListDrinks(ctx context.Context, _ *emptypb.Empty) (*pb.ListDrinksResponse, error) {
	return &pb.ListDrinksResponse{Drinks: s.drinks}, nil
}

// AddOrder implements: rpc AddOrder(Order) returns (Empty)
func (s *GRPCService) AddOrder(ctx context.Context, in *pb.Order) (*emptypb.Empty, error) {
	// map pb.Order -> model.Order
	s.orders = append(s.orders, in)
	return &emptypb.Empty{}, nil
}

// GetTotalledOrders implements: rpc GetTotalledOrders(Empty) returns (TotalledOrdersResponse)
func (s *GRPCService) GetTotalledOrders(ctx context.Context, _ *emptypb.Empty) (*pb.TotalledOrdersResponse, error) {
	totals := make(map[uint64]uint64)
	for _, order := range s.orders {
		totals[order.DrinkId] += order.Amount
	}

	return &pb.TotalledOrdersResponse{Totals: totals}, nil
}
