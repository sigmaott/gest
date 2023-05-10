package i18nfx

import (
	"github.com/gestgo/gest/package/extension/i18nfx/loader"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/fx"
)

type I18nParams struct {
	fx.In
	Loader      loader.II18nLoader   `name:"i18nLoader"`
	Translators []locales.Translator `group:"translators"`
}

func NewUniversalTranslator(
	params I18nParams,
) Result {
	enc := en.New()
	uTranslators := ut.New(enc)
	AddTranslators(uTranslators, params.Translators)
	LoadTranslate(params, uTranslators)
	return Result{
		UniversalTranslator: uTranslators,
	}
}

type Result struct {
	fx.Out
	UniversalTranslator *ut.UniversalTranslator `name:"universalTranslator"`
}

func AddTranslators(uTranslators *ut.UniversalTranslator, translators []locales.Translator) {
	for _, translator := range translators {
		err := uTranslators.AddTranslator(translator, true)
		if err != nil {
			return
		}
	}

}
func LoadTranslate(params I18nParams, uTranslators *ut.UniversalTranslator) {
	translators := params.Translators
	data := params.Loader.LoadData()
	for _, trans := range translators {

		if val, ok := data[trans.Locale()]; ok {
			transLocale, _ := uTranslators.GetTranslator(trans.Locale())

			for _, translation := range val {

				switch translation.Type {
				case "Ordinal":
					err := transLocale.AddOrdinal(translation.Key, translation.Trans, StringToPluralRule(translation.Rule), translation.Override)
					if err != nil {
						continue
					}
					continue
				case "Cardinal":
					err := transLocale.AddCardinal(translation.Key, translation.Trans, StringToPluralRule(translation.Rule), translation.Override)
					if err != nil {
						continue
					}
					continue
				case "Range":
					err := transLocale.AddRange(translation.Key, translation.Trans, StringToPluralRule(translation.Rule), translation.Override)
					if err != nil {
						continue
					}
					continue

				default:
					err := transLocale.Add(translation.Key, translation.Trans, false)
					if err != nil {
						return
					}
					continue
				}

			}

		}

	}
}
