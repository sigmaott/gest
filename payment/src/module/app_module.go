package module

import (
	"context"
	"fmt"
	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/extension/i18nfx"
	"github.com/gestgo/gest/package/extension/i18nfx/loader"
	"github.com/gestgo/gest/package/extension/logfx"
	"github.com/go-playground/locales/en"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"log"
	"os"
	"payment/config"
	"payment/src/module/payment"
)

func getCurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd

}
func BasicConnection(ctx context.Context, uri string, databaseName string) (db *mongo.Database, err error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	return client.Database(databaseName), err
}
func NewMongoConnection() *mongo.Database {
	database, err := BasicConnection(context.TODO(), config.GetConfiguration().Mongo.Uri, config.GetConfiguration().Mongo.Database)
	if err != nil {
		log.Print(err)
	}

	return database
}

func NewApp() *fx.App {

	return fx.New(
		fx.Provide(
			func() *mongo.Database {
				return NewMongoConnection()
			},
			fx.Annotate(
				echo.New,
				fx.ResultTags(`name:"platformEcho"`)),
			fx.Annotate(
				func() int {
					return config.GetConfiguration().Http.Port
				},
				fx.ResultTags(`name:"platformEchoPort"`)),
			fx.Annotate(
				SetGlobalPrefix,
				fx.ParamTags(`name:"platformEcho"`),
			),
			fx.Annotate(
				func() loader.II18nLoader {
					return loader.NewI18nJsonLoader(loader.Params{Path: fmt.Sprintf("%s/locales", getCurrentDir())})

				},
				fx.ResultTags(`name:"i18nLoader"`),
			),
			fx.Annotate(
				en.New,
				fx.ResultTags(`group:"translators"`),
			),
		),

		echofx.Module(),
		payment.Module(),
		logfx.Module(),
		i18nfx.Module(),
		fx.Invoke(EnableLogRequest),
		fx.Invoke(EnableSwagger),
		fx.Invoke(EnableErrorHandler),
		fx.Invoke(EnableNotFound),
		fx.Invoke(func(*echo.Echo) {}),
	)

}
