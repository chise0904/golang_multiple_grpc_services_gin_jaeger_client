package client

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "golang_multiple_grpc_services_gin_jaeger_client/hello"
	pb2 "golang_multiple_grpc_services_gin_jaeger_client/momo"

	"google.golang.org/grpc"
)

func Run() {

	addr := "172.17.0.3:9999"
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Can not connect to gRPC server: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.SayHello(ctx, &pb.HelloRequest{Message: "Moto"})
	if err != nil {
		log.Fatalf("Could not get nonce: %v", err)
	}

	fmt.Println("Response:", r.GetMessage())

	// ===========================================

	c2 := pb2.NewByebyeClient(conn)

	ctx2, cancel2 := context.WithTimeout(context.Background(), time.Second)
	defer cancel2()

	r2, err2 := c2.SayBye(ctx2, &pb2.MomoRequest{Message: "Byela"})
	if err != nil {
		log.Fatalf("Could not get nonce: %v", err2)
	}

	fmt.Println("Response:", r2.GetMessage())
}
