package repository

import (
	"github.com/gestgo/gest/package/core/repository"
	mongoRepository "github.com/gestgo/gest/package/technique/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"quota/src/module/quota/model"
)

type IAppQuotaRepository interface {
	repository.IRepository[model.AppQuota]
}
type appQuotaRepository struct {
	mongoRepository.BaseMongoRepository[model.AppQuota]
}

func NewAppQuotaRepository(db *mongo.Database) IAppQuotaRepository {
	return &appQuotaRepository{
		BaseMongoRepository: mongoRepository.BaseMongoRepository[model.AppQuota]{
			Collection: db.Collection("app_quotas"),
		},
	}
}
