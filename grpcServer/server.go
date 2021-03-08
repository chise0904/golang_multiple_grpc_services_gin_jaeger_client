package grpcServer

import (
	"fmt"
	"log"
	"net"
	"time"

	"golang_multiple_grpc_services_gin_jaeger_client/EchoServer"
	"golang_multiple_grpc_services_gin_jaeger_client/byeService"
	"golang_multiple_grpc_services_gin_jaeger_client/client"

	pb "golang_multiple_grpc_services_gin_jaeger_client/hello"
	pb2 "golang_multiple_grpc_services_gin_jaeger_client/momo"

	"google.golang.org/grpc"
)

func Run() {

	apiListener, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Println(err)
		return
	}

	// 註冊多個 grpc services
	es := &EchoServer.EchoServer{}
	bye := &byeService.ByeService{}

	grpc := grpc.NewServer()

	pb.RegisterGreeterServer(grpc, es)
	pb2.RegisterByebyeServer(grpc, bye)

	go runClient()

	if err := grpc.Serve(apiListener); err != nil {
		log.Fatal(" grpc.Serve Error: ", err)
		return
	}

}

func runClient() {

	time.Sleep(10 * time.Second)
	fmt.Println("After 3 seconds")

	client.Run()
}
