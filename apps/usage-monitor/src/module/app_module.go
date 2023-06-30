package module

import (
	"fmt"
	"os"
	"usage-monnitor/config"
	"usage-monnitor/src/module/usage"

	"github.com/gestgo/gest/package/extension/echofx"
	"github.com/gestgo/gest/package/extension/i18nfx"
	"github.com/gestgo/gest/package/extension/i18nfx/loader"
	"github.com/gestgo/gest/package/extension/logfx"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/vi"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
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

		echofx.ForRoot(config.GetConfiguration().Http.Port),
		fx.Provide(
			fx.Annotate(
				SetGlobalPrefix,
				fx.ParamTags(`name:"platformEcho"`),
			),
		),
		i18nfx.ForRoot(i18nfx.I18nModuleParams{
			FallbackLanguage: "en",
			Loader:           loader.NewI18nJsonLoader(loader.Params{Path: fmt.Sprintf("%s/apps/usage-monitor/locales", getCurrentDir())}),
			Translators: []locales.Translator{
				en.New(),
				vi.New(),
			},
		}),
		usage.Module(),
		logfx.Module(),
		fx.Invoke(func(*echo.Echo) {}),
	)

}
