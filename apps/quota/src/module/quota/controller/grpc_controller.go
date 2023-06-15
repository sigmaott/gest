package controller

import (
	"context"
	"github.com/gestgo/gest/package/extension/grpcfx"
	"github.com/samber/lo"
	"google.golang.org/grpc"
	"quota/src/module/quota/model"
	"quota/src/module/quota/service"
	pb "sigma-streaming/service/gen"
)

type IQuotaGrpcController interface {
	grpcfx.IGrpcController
	pb.QuotaServiceServer
}
type quotaGrpcController struct {
	service service.IQuotaService
}

func (q *quotaGrpcController) GetQuotaByAppId(ctx context.Context, request *pb.GetQuotaByAppIdRequest) (*pb.GetQuotaByAppIdResponse, error) {
	res, err := q.service.GetAppQuota(request.AppId)
	if err != nil {
		return nil, err
	}
	quotas := lo.MapValues(res.Quotas, func(x model.Quota, _ string) *pb.Quota {
		return &pb.Quota{
			Resource: x.Resource,
			Hard:     x.Hard,
		}
	})
	return &pb.GetQuotaByAppIdResponse{
		AppId:      res.AppId,
		QuotaGroup: res.GroupQuota,
		Quotas:     quotas,
	}, nil
}

func (q *quotaGrpcController) GetQuotaResourceByAppId(ctx context.Context, request *pb.GetQuotaResourceByAppIdRequest) (*pb.GetQuotaResourceByAppIdResponse, error) {
	res, err := q.service.GetResourceAppQuota(request.AppId, request.Resource)
	if err != nil {
		return nil, err
	}

	return &pb.GetQuotaResourceByAppIdResponse{
		Quota: &pb.Quota{
			Resource: res.Resource,
			Hard:     res.Hard,
		},
	}, nil
}

func (q *quotaGrpcController) UpsertQuotaByAppId(ctx context.Context, request *pb.UpsertQuotaByAppIdRequest) (*pb.UpsertQuotaByAppIdResponse, error) {
	err := q.service.UpsertAppQuota(request.AppId, request.QuotaGroup)
	if err != nil {
		return nil, err
	}
	return &pb.UpsertQuotaByAppIdResponse{
		Success: true,
	}, nil
}

func (q *quotaGrpcController) RegisterGrpcController(server *grpc.Server) {
	pb.RegisterQuotaServiceServer(server, q)
}

func NewQuotaGrpcController(service service.IQuotaService) IQuotaGrpcController {
	return &quotaGrpcController{
		service: service,
	}
}
