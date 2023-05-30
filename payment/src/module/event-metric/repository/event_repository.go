package repository

import (
	"context"
	"errors"
	"fmt"
	"github.com/getlago/lago-go-client"
	"go.uber.org/zap"
	"log"
	"time"
)

type IEventRepository interface {
	CreateEvent(code string, properties map[string]any, appId string) error
}
type eventRepository struct {
	client *lago.Client
	logger *zap.SugaredLogger
}

func (i *eventRepository) CreateEvent(code string, properties map[string]any, appId string) error {
	log.Print(&lago.EventInput{
		TransactionID:      fmt.Sprintf("%s_%s_%d", code, appId, time.Now().Unix()),
		ExternalCustomerID: appId,
		//ExternalSubscriptionID: "giangnt_sb",
		Code:       code,
		Timestamp:  time.Now().Unix(),
		Properties: properties,
	})
	err := i.client.Event().Create(context.TODO(), &lago.EventInput{
		TransactionID:      fmt.Sprintf("%s_%s_%d", code, appId, time.Now().Unix()),
		ExternalCustomerID: appId,
		Code:               code,
		Timestamp:          time.Now().Unix(),
		Properties:         properties,
	})
	if err != nil {

		return errors.New(fmt.Sprintf("%+v", err))
	}
	return nil

}

func NewEventRepository(client *lago.Client, logger *zap.SugaredLogger) IEventRepository {
	return &eventRepository{
		client: client,
		logger: logger,
	}

}
