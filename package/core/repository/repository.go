package repository

import "context"

type Sort map[string]string

type Paginate struct {
	Offset int64
	Limit  int64
}

const (
	DESC = "desc"
	ASC  = "asc"
)

type IRepository[T any] interface {
	Upsert(ctx context.Context, query any, data *T) (result *T, err error)
	CreateOne(ctx context.Context, data *T) (result *T, err error)
	FindOne(ctx context.Context, query any) (result *T, err error)
	UpdateOne(ctx context.Context, query any, data any) (result *T, err error)
	DeleteOne(ctx context.Context, query any) (result *T, err error)
	CreateMany(ctx context.Context, data []any) (err error)
	UpdateMany(ctx context.Context, query any, data any) (err error)
	DeleteMany(ctx context.Context, query any) (err error)
	Count(ctx context.Context, query any) (count int64, err error)
	FindAll(ctx context.Context, query any, paginate *Paginate, sort *Sort) (results []*T, err error)
}
