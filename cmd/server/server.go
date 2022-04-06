package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/plummertr/users"
	pb "github.com/plummertr/users/api/internal-api/v1"
	"google.golang.org/grpc"
)

func main() {
	var port string
	flag.StringVar(&port, "port", "", "Port to run server on")

	flag.Parse()

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v\n", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterUserServiceServer(grpcServer, users.NewUserServer())

	fmt.Print("running server on ", port)

	grpcServer.Serve(lis)
}
