package i18nfx

import (
	"github.com/gestgo/gest/package/extension/i18nfx/loader"
	"github.com/go-playground/locales"
	"go.uber.org/fx"
)

type I18nModuleParams struct {
	FallbackLanguage string
	Loader           loader.II18nLoader
	Translators      []locales.Translator
}

func ForRoot(params I18nModuleParams) fx.Option {
	return fx.Module("i18nfx",
		fx.Provide(fx.Annotate(
			func() []locales.Translator {
				return params.Translators
			},
			fx.ResultTags(`name:"translators"`),
		)),
		fx.Provide(fx.Annotate(
			func() loader.II18nLoader {
				return params.Loader
			},
			fx.ResultTags(`name:"i18nLoader"`),
		)),
		fx.Provide(fx.Annotate(
			func() string {
				return params.FallbackLanguage
			},
			fx.ResultTags(`name:"fallbackLanguage"`),
		)),
		fx.Provide(newUniversalTranslator, NewI18nService))
}
