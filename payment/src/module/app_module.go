package module

import (
	"context"
	"fmt"
	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/extension/i18nfx"
	"github.com/gestgo/gest/package/extension/i18nfx/loader"
	"github.com/gestgo/gest/package/extension/kafkafx"
	"github.com/gestgo/gest/package/extension/logfx"
	"github.com/getlago/lago-go-client"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/vi"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
	"os"
	"payment/config"
	event_metric "payment/src/module/event-metric"
	"payment/src/module/health"
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

//	func NewMongoConnection() *mongo.Database {
//		database, err := BasicConnection(context.TODO(), config.GetConfiguration().Mongo.Uri, config.GetConfiguration().Mongo.Database)
//		if err != nil {
//			log.Print(err)
//		}
//
//		return database
//	}
func NewLagoClient() *lago.Client {
	client := lago.New().
		SetBaseURL(fmt.Sprintf("http://%s:%s", config.GetConfiguration().Lago.Host, config.GetConfiguration().Lago.Port)).
		SetDebug(true).
		SetApiKey(config.GetConfiguration().Lago.ApiKey)
	return client
}
func NewApp() *fx.App {

	return fx.New(
		fx.Provide(

			fx.Annotate(
				func() *kafkafx.KafkaSubscriber {
					return &kafkafx.KafkaSubscriber{}
				},
				fx.ResultTags(`name:"platformKafka"`)),
			NewLagoClient,
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
		fx.Provide(
			func() *validator.Validate {
				return validator.New()
			},
			fx.Annotate(
				func(ut *ut.UniversalTranslator) *ut.UniversalTranslator {
					english := en.New()
					vietnam := vi.New()
					ut.AddTranslator(english, false)
					ut.AddTranslator(vietnam, false)
					return ut
				},
				fx.ParamTags(`name:"universalTranslator"`),
			),

			NewI18nValidate,
		),
		echofx.Module(),
		event_metric.Module(),
		logfx.Module(),
		i18nfx.Module(),
		kafkafx.Module(),
		health.Module(),
		fx.Invoke(RegisterValidateTranslations),
		fx.Invoke(EnableValidationRequest),
		fx.Invoke(EnableLogRequest),
		fx.Invoke(EnableErrorHandler),
		fx.Invoke(EnableNotFound),
		//fx.Invoke(EnableSwagger),
		fx.Invoke(EnableLogRouter),
		fx.Invoke(func(*echo.Echo) {}),
		fx.Invoke(func(*kafkafx.KafkaSubscriber) {}),
	)

}
