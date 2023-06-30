package repository

import (
	"context"
	pb "sigma-streaming/service/gen"
)

type IQuotaRepository interface {
	GetQuotaResource(appId string, resource string) (*pb.GetQuotaResourceByAppIdResponse, error)
}

type quotaRepository struct {
	client pb.QuotaServiceClient
}

func (q *quotaRepository) GetQuotaResource(appId string, resource string) (*pb.GetQuotaResourceByAppIdResponse, error) {
	return q.client.GetQuotaResourceByAppId(context.TODO(), &pb.GetQuotaResourceByAppIdRequest{
		AppId:    appId,
		Resource: resource,
	})
}

func NewIQuotaRepository(client pb.QuotaServiceClient) IQuotaRepository {
	return &quotaRepository{
		client: client,
	}

}
