package module

import (
	"context"
	"fmt"
	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/extension/grpcfx"
	"github.com/gestgo/gest/package/extension/i18nfx"
	"github.com/gestgo/gest/package/extension/i18nfx/loader"
	"github.com/gestgo/gest/package/extension/kafkafx"
	"github.com/gestgo/gest/package/extension/logfx"
	"github.com/gestgo/gest/package/extension/mongofx"
	"github.com/go-playground/locales/en"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"os"
	"quota/config"
	"quota/src/module/health"
	"quota/src/module/quota"
)

func getCurrentDir() string {
	pwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return pwd

}

func NewApp() *fx.App {

	return fx.New(
		fx.Provide(

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

		logfx.Module(),
		i18nfx.Module(),
		kafkafx.Module(),
		grpcfx.ForRoot(fmt.Sprintf("%s:%d", config.GetConfiguration().Grpc.Host, config.GetConfiguration().Grpc.Port)),
		mongofx.ForRoot(context.TODO(), config.GetConfiguration().Mongo.Uri, config.GetConfiguration().Mongo.Database),
		quota.Module(),
		health.Module(),
		//fx.Invoke(EnableSwagger),
		fx.Invoke(func(*echo.Echo) {}),
		fx.Invoke(func(server *grpc.Server) {}),
	)

}
