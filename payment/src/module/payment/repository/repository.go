package repository

import (
	"github.com/gestgo/gest/package/core/repository"
	mongoRepository "github.com/gestgo/gest/package/technique/mongo"
	"go.mongodb.org/mongo-driver/mongo"
	"payment/src/module/payment/model"
)

type IPaymentRepository interface {
	repository.IRepository[model.Payment]
}
type paymentRepository struct {
	mongoRepository.BaseMongoRepository[model.Payment]
}

func NewPaymentRepository(db *mongo.Database) IPaymentRepository {
	return &paymentRepository{
		BaseMongoRepository: mongoRepository.BaseMongoRepository[model.Payment]{
			Collection: db.Collection("payment"),
		},
	}

}
