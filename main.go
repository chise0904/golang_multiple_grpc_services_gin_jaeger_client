package main

import (
	// "golang_multiple_grpc_services_gin_jaeger_client/grpcServer"
	// "golang_multiple_grpc_services_gin_jaeger_client/httpServer"
	"golang_multiple_grpc_services_gin_jaeger_client/client"
)

func main() {

	ch := make(chan struct{})

	// go grpcServer.Run()
	// go httpServer.Run()

	client.Run()

	<-ch
}
