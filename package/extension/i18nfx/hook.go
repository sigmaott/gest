package i18nfx

import (
	"log"

	"github.com/gestgo/gest/package/extension/i18nfx/loader"
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/fx"
)

type I18nParams struct {
	fx.In
	Loader      loader.II18nLoader   `name:"i18nLoader"`
	Translators []locales.Translator `name:"translators"`
}

func newUniversalTranslator(
	params I18nParams,
) Result {
	log.Print(params.Translators)
	enc := en.New()
	uTranslators := ut.New(enc)
	err := addTranslators(uTranslators, params.Translators)
	if err != nil {
		log.Fatal(err)
	}

	uTranslators = loadTranslate(params, uTranslators)
	return Result{
		UniversalTranslator: uTranslators,
	}
}

type Result struct {
	fx.Out
	UniversalTranslator *ut.UniversalTranslator `name:"universalTranslator"`
}

func addTranslators(uTranslators *ut.UniversalTranslator, translators []locales.Translator) error {
	for _, translator := range translators {
		err := uTranslators.AddTranslator(translator, true)
		if err != nil {
			return err
		}
	}
	return nil

}
func loadTranslate(params I18nParams, uTranslators *ut.UniversalTranslator) *ut.UniversalTranslator {
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
						return uTranslators
					}
					continue
				}

			}

		}

	}
	return uTranslators
}
