package i18nfx

import "go.uber.org/fx"

func Module() fx.Option {
	return fx.Module("i18nfx", fx.Provide(NewUniversalTranslator, NewI18nService))
}
