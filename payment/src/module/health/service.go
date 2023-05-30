package health

import (
	"context"
	"errors"
	healthcheck "github.com/gestgo/gest/technique/health"
	"github.com/labstack/gommon/log"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"payment/config"
	"time"
)

type IHeathCheckService interface {
	HeathCheck() (*healthcheck.Response, error)
}

type heathCheckService struct {
	logger *zap.SugaredLogger
}

func (h *heathCheckService) HeathCheck() (*healthcheck.Response, error) {
	res := healthcheck.HandlerHeathCheck(healthcheck.WithTimeout(5*time.Second),
		healthcheck.WithChecker("kafka", healthcheck.CheckerFunc(
			func(ctx context.Context) error {
				var errs []error
				for _, url := range config.GetConfiguration().Kafka.Urls {
					_, err := kafka.DialContext(ctx, "tcp", url)
					if err != nil {
						errs = append(errs, err)
					}
				}
				log.Print("run kafka ")
				if len(errs) == len(config.GetConfiguration().Kafka.Urls) {
					var errWarap error
					for _, err := range errs {
						if errWarap == nil {
							errWarap = err
						} else {
							errWarap = err
						}

					}

					return errWarap
				}
				return nil
			},
		)),
	)
	log.Print(res.Errors)
	if len(res.Errors) > 0 {
		return res, errors.New("Service Unavailable")
	}
	return res, nil
}

func NewHeathCheckService(logger *zap.SugaredLogger) IHeathCheckService {
	return &heathCheckService{
		logger: logger,
	}
}
