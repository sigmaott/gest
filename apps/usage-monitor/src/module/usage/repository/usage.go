package repository

import (
	"context"
	"errors"
	"strconv"
	"usage-monnitor/src/module/usage/model"

	"github.com/getlago/lago-go-client"
	"github.com/samber/lo"
)

type IUsageRepository interface {
	GetUsageByMetric(appId string, metricCode string) (*model.Usage, error)
}
type usageRepository struct {
	client lago.Client
}

func (u *usageRepository) GetUsageByMetric(appId string, metricCode string) (*model.Usage, error) {
	subs, err := u.client.Subscription().GetList(context.TODO(), lago.SubscriptionListInput{
		ExternalCustomerID: appId,
	})
	if err != nil {
		return nil, err.Err
	}

	//Only 1 subscription
	if len(subs.Subscriptions) == 0 {
		return nil, errors.New("No Usage Found")
	}
	cuUsage, err := u.client.Customer().CurrentUsage(context.TODO(), appId, &lago.CustomerUsageInput{
		ExternalSubscriptionID: subs.Subscriptions[0].ExternalID,
	})
	if err != nil {
		return nil, err.Err
	}

	metric, ok := lo.Find(cuUsage.ChargesUsage, func(item lago.CustomerChargeUsage) bool {
		return item.BillableMetric.Code == metricCode
	})
	if !ok {
		return &model.Usage{
			Code:  metricCode,
			Usage: 0,
		}, nil
	}
	unit, errParse := strconv.ParseFloat(metric.Units, 32)
	if errParse != nil {
		return &model.Usage{
			Code:  metricCode,
			Usage: 0,
		}, nil
	}

	return &model.Usage{
		Code:  metricCode,
		Usage: unit,
	}, nil
}

func NewUsageRepository(client lago.Client) IUsageRepository {
	return &usageRepository{
		client: client,
	}
}
