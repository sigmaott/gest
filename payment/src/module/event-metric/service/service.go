package service

import (
	"github.com/gestgo/gest/package/extension/i18nfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"log"
	"payment/config"
	"payment/src/module/event-metric/dto"
	"payment/src/module/event-metric/model"
	"payment/src/module/event-metric/repository"
)

type ISSAIEventService interface {
	Create(event *dto.CreateEvent) error
}
type ssaiEventService struct {
	ssaiEventRepository repository.IEventRepository
	logger              *zap.SugaredLogger
}

func (u *ssaiEventService) Create(event *dto.CreateEvent) error {
	log.Print(config.GetConfiguration().Lago.BillableMetric)
	model := model.SSAIAdsInsertProperties{TotalAdsInsert: event.TotalAdsInsert}
	err := u.ssaiEventRepository.CreateEvent(config.GetConfiguration().Lago.BillableMetric.SSAIInsertAdsCode, model.StructToMap(), event.AppId)
	if err != nil {
		return err
	}
	return err
}

type UserServiceParams struct {
	fx.In
	I18nService i18nfx.I18nService
}

func NewEventService(ssaiEventRepository repository.IEventRepository, logger *zap.SugaredLogger) ISSAIEventService {

	return &ssaiEventService{
		ssaiEventRepository: ssaiEventRepository,
		logger:              logger,
	}

}
