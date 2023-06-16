package health

import (
	"context"
	"errors"
	healthcheck "github.com/gestgo/gest/technique/health"
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
	"time"
)

type IHeathCheckService interface {
	HeathCheck() (*healthcheck.Response, error)
}

type heathCheckService struct {
	logger   *zap.SugaredLogger
	database *mongo.Database
}

func (h *heathCheckService) HeathCheck() (*healthcheck.Response, error) {
	res := healthcheck.HandlerHeathCheck(healthcheck.WithTimeout(5*time.Second),
		healthcheck.WithChecker("mongodb", healthcheck.CheckerFunc(
			func(ctx context.Context) error {
				err := h.database.Client().Ping(ctx, nil)
				if err != nil {
					h.logger.Info(err)
				}
				return err
			},
		)),
	)
	log.Print(res.Errors)
	if len(res.Errors) > 0 {
		return res, errors.New("Service Unavailable")
	}
	return res, nil
}

func NewHeathCheckService(logger *zap.SugaredLogger, database *mongo.Database) IHeathCheckService {
	return &heathCheckService{
		logger:   logger,
		database: database,
	}
}
