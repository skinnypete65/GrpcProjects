package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	pb "ordermanagment/server/proto/order_management"
)

type server struct {
	orderMap map[string][]*pb.Order
	pb.UnimplementedOrderManagementServer
}

func newServer() *server {
	return &server{orderMap: make(map[string][]*pb.Order)}
}

func (s *server) SearchOrders(id *pb.ClientID, stream pb.OrderManagement_SearchOrdersServer) error {
	for _, order := range s.orderMap[id.ClientID] {
		err := stream.Send(order)
		if err != nil {
			return fmt.Errorf("error sending message to stream : %v", err)
		}
	}

	return nil
}

func (s *server) AddOrder(ctx context.Context, order *pb.ClientOrder) (*emptypb.Empty, error) {
	clientID := order.ClientID
	s.orderMap[clientID] = append(s.orderMap[clientID], order.Order)

	return &emptypb.Empty{}, status.New(codes.OK, "Order has been added").Err()
}
