package gorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/gestgo/gest/package/core/repository"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type BaseGormRepository[T any] struct {
	Db    *gorm.DB
	Table string
}

func (b *BaseGormRepository[T]) FindAll(ctx context.Context, query any, paginate *repository.Paginate, sort *repository.Sort) (results []*T, err error) {
	results = []*T{}
	querySQL := b.Db.Table(b.Table)
	if paginate != nil {
		querySQL = querySQL.Offset(int(paginate.Offset)).Limit(int(paginate.Limit))
	}

	if sort != nil {

		for k, v := range *sort {
			var sortQuery string
			if v == repository.DESC {
				sortQuery = fmt.Sprintf("%s DESC", k)
			}
			if v == repository.ASC {
				sortQuery = fmt.Sprintf("%s ASC", k)
			}
			querySQL = querySQL.Order(sortQuery)

		}

	}
	results = []*T{}
	err = b.Db.Table(b.Table).Where(query).Find(&results).Error

	return results, err

}

func (b *BaseGormRepository[T]) CreateOne(ctx context.Context, data *T) (result *T, err error) {
	err = b.Db.Table(b.Table).Create(data).Error
	return data, err
}
func (b *BaseGormRepository[T]) FindOne(ctx context.Context, query any) (result *T, err error) {
	result = new(T)
	err = b.Db.Table(b.Table).Where(query).Take(&result).Error
	return
}

func (b *BaseGormRepository[T]) UpdateOne(ctx context.Context, query any, data any) (result *T, err error) {
	err = b.Db.Table(b.Table).Where(query).Updates(data).Error
	return
}

func (b *BaseGormRepository[T]) DeleteOne(ctx context.Context, query any) (result *T, err error) {
	result = new(T)
	err = b.Db.Table(b.Table).Where(query).Delete(result).Error
	return result, err
}

func (b *BaseGormRepository[T]) UpdateMany(ctx context.Context, query any, data any) (err error) {
	err = b.Db.Table(b.Table).Where(query).Updates(data).Error
	return err
}

func (b *BaseGormRepository[T]) DeleteMany(ctx context.Context, query any) (err error) {
	model := new(T)
	err = b.Db.Table(b.Table).Where(query).Delete(model).Error
	return err
}

func (b *BaseGormRepository[T]) Upsert(ctx context.Context, query any, data *T) (result *T, err error) {
	result = new(T)
	err = b.Db.Table(b.Table).Where(query).First(result).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		} else {
			//insert
			if err = b.Db.Table(b.Table).Where(query).Create(data).Error; err != nil {
				return nil, err
			}
		}
	}
	if err = b.Db.Table(b.Table).Where(query).Save(data).Error; err != nil {
		log.Print(err)
		return nil, err
	}
	return data, err
}

func (b *BaseGormRepository[T]) CreateMany(ctx context.Context, data []any) (err error) {

	err = b.Db.Table(b.Table).CreateInBatches(data, len(data)).Error
	return err
}

func (b *BaseGormRepository[T]) Count(ctx context.Context, query any) (count int64, err error) {
	err = b.Db.Table(b.Table).Where(query).Count(&count).Error
	return
}

func NewBaseOrmRepository[T any](db *gorm.DB, table string) *BaseGormRepository[T] {
	return &BaseGormRepository[T]{
		Db: db, Table: table,
	}
}

func BuildQuery(query interface{}, args ...interface{}) clause.Where {
	state := new(gorm.Statement)
	return clause.Where{Exprs: state.BuildCondition(query, args...)}
}
