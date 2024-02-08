package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	pb "ordermanagment/client/proto/order_management"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	searchStream, _ := c.SearchOrders(ctx, &pb.ClientID{ClientID: "1"})

	for {
		searchOrder, err := searchStream.Recv()
		if err != nil {
			break
		}
		log.Print("Search Result : ", searchOrder)
	}
}
