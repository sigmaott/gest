package module

import (
	"context"
	"fmt"
	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/extension/grpcfx"
	"github.com/gestgo/gest/package/extension/i18nfx"
	"github.com/gestgo/gest/package/extension/logfx"
	"github.com/gestgo/gest/package/extension/mongofx"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"google.golang.org/grpc"
	"usage-monnitor/config"
	"usage-monnitor/src/module/health"
)

func NewApp() *fx.App {

	return fx.New(
		fx.Provide(
			fx.Annotate(
				SetGlobalPrefix,
				fx.ParamTags(`name:"platformEcho"`),
			),
		),

		echofx.ForRoot(config.GetConfiguration().Http.Port),
		grpcfx.ForRoot(fmt.Sprintf("%s:%d", config.GetConfiguration().Grpc.Host, config.GetConfiguration().Grpc.Port)),
		mongofx.ForRoot(context.TODO(), config.GetConfiguration().Mongo.Uri, config.GetConfiguration().Mongo.Database),
		logfx.Module(),
		i18nfx.Module(),
		usage.Module(),
		health.Module(),
		fx.Invoke(func(*echo.Echo) {}),
		fx.Invoke(func(server *grpc.Server) {}),
	)

}
