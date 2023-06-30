package controller

import (
	"context"

	"github.com/gestgo/gest/package/extension/grpcfx"
	"google.golang.org/grpc"

	pb "sigma-streaming/service/gen"
)

type ISSAIUsageMonitorGrpcController interface {
	grpcfx.IGrpcController
	pb.SSAIUsageMonitorServiceServer
}
type ssaiUsageMonitorGrpcController struct {
}

func (q *ssaiUsageMonitorGrpcController) CheckQuota(ctx context.Context, request *pb.CheckQuotaRequest) (*pb.CheckQuotaResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (q *ssaiUsageMonitorGrpcController) RegisterGrpcController(server *grpc.Server) {
	pb.RegisterSSAIUsageMonitorServiceServer(server, q)
}

func NewQuotaGrpcController() ISSAIUsageMonitorGrpcController {
	return &ssaiUsageMonitorGrpcController{}
}
