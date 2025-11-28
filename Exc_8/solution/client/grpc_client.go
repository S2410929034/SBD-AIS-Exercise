package client

import (
	"context"
	"exc8/pb"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GrpcClient struct {
	client pb.OrderServiceClient
	conn   *grpc.ClientConn
}

func NewGrpcClient() (*GrpcClient, error) {
	conn, err := grpc.NewClient(":4000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := pb.NewOrderServiceClient(conn)
	return &GrpcClient{client: client}, nil
}

func (c *GrpcClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

func (c *GrpcClient) Run() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// todo
	// 1. List drinks
	fmt.Println("Requesting drinks 🍹🍺☕")
	drinksResp, err := c.client.ListDrinks(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("ListDrinks failed: %w", err)
	}

	fmt.Println("Available drinks:")
	for _, d := range drinksResp.Drinks {
		fmt.Printf("\t> id:%d  name:%q  price:%.0f  description:%q\n", d.Id, d.Name, d.Price, d.Description)
	}

	if len(drinksResp.Drinks) == 0 {
		return fmt.Errorf("no drinks available")
	}

	// 2. Order a few drinks
	fmt.Println("Ordering drinks 👨‍🍳⏱️🍻🍻")
	for _, drink := range drinksResp.Drinks {
		fmt.Printf("\t> Ordering: 2 x %s\n", drink.Name)
		_, err = c.client.AddOrder(ctx, &pb.Order{
			DrinkId: drink.Id,
			Amount:  2,
		})
		if err != nil {
			return fmt.Errorf("AddOrder failed: %w", err)
		}
	}

	// 3. Order more drinks
	fmt.Println("Ordering another round of drinks 👨‍🍳⏱️🍻🍻")
	for _, drink := range drinksResp.Drinks {
		fmt.Printf("\t> Ordering: 6 x %s\n", drink.Name)
		_, err = c.client.AddOrder(ctx, &pb.Order{
			DrinkId: drink.Id,
			Amount:  6,
		})
		if err != nil {
			return fmt.Errorf("AddOrder (second round) failed: %w", err)
		}
	}

	// 4. Get order total
	fmt.Println("Getting the bill 💹💹💹")
	totalResp, err := c.client.GetTotalledOrders(ctx, &emptypb.Empty{})
	if err != nil {
		return fmt.Errorf("GetTotalledOrders failed: %w", err)
	}

	// Map drink IDs to names for display
	drinkNames := make(map[uint64]string)
	for _, d := range drinksResp.Drinks {
		drinkNames[d.Id] = d.Name
	}

	for drinkID, total := range totalResp.Totals {
		drinkName := drinkNames[drinkID]
		fmt.Printf("\t> Total: %d x %s\n", total, drinkName)
	}

	// print responses after each call
	return nil
}
