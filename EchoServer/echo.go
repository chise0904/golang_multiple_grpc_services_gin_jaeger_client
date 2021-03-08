package EchoServer

import (
	context "context"
	"fmt"
	pb "golang_multiple_grpc_services_gin_jaeger_client/hello"
)

type EchoServer struct{}

func (e *EchoServer) SayHello(ctx context.Context, req *pb.HelloRequest) (resp *pb.HelloReply, err error) {

	fmt.Println("[Server receive client request]" + req.GetMessage())
	return &pb.HelloReply{
		Message: "[Echo From Server] " + req.GetMessage(),
	}, nil

}
