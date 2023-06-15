package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.uber.org/zap"
	"quota/src/module/quota/model"
	"quota/src/module/quota/repository"
	"time"
)

type IQuotaService interface {
	GetAppQuota(appId string) (*model.AppQuota, error)
	GetResourceAppQuota(appId string, resource string) (*model.Quota, error)
	UpsertAppQuota(appId string, groupQuotaName string) error
}
type quotaService struct {
	baseQuotaRepository repository.BaseQuotaRepository
	appQuotaRepository  repository.IAppQuotaRepository
	logger              *zap.SugaredLogger
}

func (q *quotaService) GetResourceAppQuota(appId string, resource string) (*model.Quota, error) {
	appQuota, err := q.GetAppQuota(appId)
	if err != nil {
		q.logger.Errorf("%+v", err)
		return nil, err
	}
	quota, ok := appQuota.Quotas[resource]
	if !ok {
		return nil, errors.New("resource not found")
	}
	return &quota, err
}

func (q *quotaService) UpsertAppQuota(appId string, groupQuotaName string) error {
	quota, err := q.baseQuotaRepository.GetBaseQuota(groupQuotaName)
	if err != nil {
		q.logger.Errorf("%+v", err)
		return err
	}

	appQuota := &model.AppQuota{
		ID:         uuid.NewString(),
		AppId:      appId,
		GroupQuota: quota.Name,
		CreatedAt:  time.Now(),
		UpdateAt:   time.Now(),
	}

	_, err = q.appQuotaRepository.Upsert(context.TODO(), bson.M{
		"appId": appId,
	}, appQuota)
	return err

}

func (q *quotaService) GetAppQuota(appId string) (*model.AppQuota, error) {

	appQuota, err := q.appQuotaRepository.FindOne(context.TODO(), bson.M{"appId": appId})
	if err != nil {
		q.logger.Errorf("%+v", err)
		return nil, err
	}
	quota, err := q.baseQuotaRepository.GetBaseQuota(appQuota.GroupQuota)
	if err != nil {
		q.logger.Errorf("%+v", err)
		return nil, err
	}
	appQuota.Quotas = quota.Quotas
	return appQuota, nil
}

func NewQuotaService(baseQuotaRepository repository.BaseQuotaRepository, appQuotaRepository repository.IAppQuotaRepository, logger *zap.SugaredLogger) IQuotaService {
	return &quotaService{
		baseQuotaRepository: baseQuotaRepository,
		appQuotaRepository:  appQuotaRepository,
		logger:              logger,
	}
}
