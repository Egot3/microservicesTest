package internal

import (
	"context"
	"fmt"
	"time"

	pb "github.com/Egot3/microservicesTest/proto/gen"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func main() bool {
	conn, err := grpc.NewClient("localhost:50051",
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(fmt.Errorf("Couldn't connect to gRPC: %v", err.Error()))
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		getStatus, err := client.HealthCheck(ctx, &emptypb.Empty{})
		if err == nil && getStatus.Alive {
			fmt.Printf("Server's health check succeded (att %d/5)", i+1)
		}
		fmt.Printf("Server's health check failed (att %d/5): %v", i+1, err)
	}

	return false
}
