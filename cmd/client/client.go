package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/plummertr/users/api/internal-api/v1"
	"google.golang.org/grpc"
)

func main() {
	var addr string

	flag.StringVar(&addr, "addr", "", "Address of gRPC server")
	flag.Parse()

	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dail: ", err)
	}
	defer c.Close()

	client := pb.NewUserServiceClient(c)

	ctx := context.Background()

	req := pb.ReadUserRequest{
		Id: int64(1),
	}

	x, err := client.ReadUser(ctx, &req)
	if err != nil {
		log.Fatalf("failed to read user", err)
	}
	fmt.Println(x)
}
