package byeService

import (
	context "context"
	"fmt"
	pb "golang_multiple_grpc_services_gin_jaeger_client/momo"
)

type ByeService struct{}

func (e *ByeService) SayBye(ctx context.Context, req *pb.MomoRequest) (resp *pb.MomoReply, err error) {

	fmt.Println("[ByeService receive client request]" + req.GetMessage())
	return &pb.MomoReply{
		Message: "[Echo From ByeService] " + req.GetMessage(),
	}, nil
}
