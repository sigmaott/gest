package mongo

import (
	"context"

	"github.com/sigmaott/gest/package/technique/database/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type BaseMongoRepository[T any] struct {
	Collection *mongo.Collection
}

func (b *BaseMongoRepository[T]) FindAll(ctx context.Context, query any, paginate *repository.Paginate, sort *repository.Sort) (results []*T, err error) {
	opt := options.Find()
	results = []*T{}
	if paginate != nil {
		opt.SetLimit(paginate.Limit).SetSkip(paginate.Offset)
	}
	if sort != nil {
		sortMongo := bson.M{}
		for k, v := range *sort {
			if v == repository.DESC {
				sortMongo[k] = -1
			}
			if v == repository.ASC {
				sortMongo[k] = 1
			}

		}
		opt.SetSort(sortMongo)
	}
	cur, err := b.Collection.Find(ctx, query, opt)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {
		//Create a value into which the single document can be decoded
		elem := new(T)
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)

	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return results, nil

}

func (b *BaseMongoRepository[T]) CreateOne(ctx context.Context, data *T) (result *T, err error) {
	_, err = b.Collection.InsertOne(ctx, data)
	return data, err
}
func (b *BaseMongoRepository[T]) FindOne(ctx context.Context, query any) (result *T, err error) {
	result = new(T)
	err = b.Collection.FindOne(ctx, query).Decode(result)
	return
}

func (b *BaseMongoRepository[T]) UpdateOne(ctx context.Context, query any, data any) (result *T, err error) {
	result = new(T)
	opt := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	err = b.Collection.FindOneAndUpdate(ctx, query, bson.M{
		"$set": data,
	}, opt).Decode(result)
	return
}

func (b *BaseMongoRepository[T]) DeleteOne(ctx context.Context, query any) (result *T, err error) {
	result = new(T)
	opt := options.FindOneAndDelete()
	err = b.Collection.FindOneAndDelete(ctx, query, opt).Decode(result)
	return
}

func (b *BaseMongoRepository[T]) UpdateMany(ctx context.Context, query any, data any) (err error) {
	_, err = b.Collection.UpdateMany(ctx, query, bson.M{
		"$set": data,
	})
	return err
}

func (b *BaseMongoRepository[T]) DeleteMany(ctx context.Context, query any) (err error) {
	_, err = b.Collection.DeleteMany(ctx, query)
	return err
}

func (b *BaseMongoRepository[T]) Upsert(ctx context.Context, query any, data *T) (result *T, err error) {
	result = new(T)
	opt := options.FindOneAndUpdate().SetUpsert(true).SetReturnDocument(options.After)
	err = b.Collection.FindOneAndUpdate(ctx, query, bson.M{
		"$set": data,
	}, opt).Decode(result)
	return
}

func (b *BaseMongoRepository[T]) CreateMany(ctx context.Context, data []any) (err error) {
	_, err = b.Collection.InsertMany(ctx, data)
	return
}

func (b *BaseMongoRepository[T]) Count(ctx context.Context, query any) (count int64, err error) {
	count, err = b.Collection.CountDocuments(ctx, query)
	return
}

func (b *BaseMongoRepository[T]) Paginate(ctx context.Context, query any, paginate *repository.Paginate, sort *repository.Sort) (results *repository.PaginateResponse[T], err error) {
	res := new(repository.PaginateResponse[T])
	res.Data = []*T{}
	count, err := b.Count(ctx, query)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return res, nil
	}
	if paginate.Limit <= 0 {
		paginate.Limit = 10
	}
	data, err := b.FindAll(ctx, query, paginate, sort)
	if err != nil {
		return nil, err
	}
	res.Page = (paginate.Offset / paginate.Limit) + 1
	res.PerPage = paginate.Limit
	res.Data = data
	res.Total = count

	return res, nil
}

func NewBaseRepository[T any](db *mongo.Database, collectionName string) repository.IRepository[T] {
	return &BaseMongoRepository[T]{
		Collection: db.Collection(collectionName),
	}
}
