package i18nfx

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/fx"
)

type II18nService interface {
	T(lang string, key string, params ...string) (string, error)

	C(lang string, key string, num float64, digits uint64, param string) (string, error)

	O(lang string, key string, num float64, digits uint64, param string) (string, error)

	R(lang string, key string, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) (string, error)
}
type I18nService struct {
	i18n *ut.UniversalTranslator
}

func (i *I18nService) C(lang string, key string, num float64, digits uint64, param string) (string, error) {
	trans, found := i.i18n.GetTranslator(lang)
	if !found {
		trans = i.i18n.GetFallback()
	}
	if message, err := trans.C(key, num, digits, param); err != nil {
		return key, err
	} else {
		return message, err
	}
}

func (i *I18nService) O(lang string, key string, num float64, digits uint64, param string) (string, error) {
	trans, found := i.i18n.GetTranslator(lang)
	if !found {
		trans = i.i18n.GetFallback()
	}
	if message, err := trans.O(key, num, digits, param); err != nil {
		return key, err
	} else {
		return message, err
	}
}

func (i *I18nService) R(lang string, key string, num1 float64, digits1 uint64, num2 float64, digits2 uint64, param1, param2 string) (string, error) {
	trans, found := i.i18n.GetTranslator(lang)
	if !found {
		trans = i.i18n.GetFallback()
	}
	if message, err := trans.R(key, num1, digits1, num2, digits2, param1, param2); err != nil {
		return key, err
	} else {
		return message, err
	}
}

func (i *I18nService) T(lang string, key string, params ...string) (string, error) {
	trans, found := i.i18n.GetTranslator(lang)
	if !found {
		trans = i.i18n.GetFallback()
	}
	//b, _ := json.Marshal(trans)

	if message, err := trans.T(key, params...); err != nil {
		return key, err
	} else {
		return message, err
	}

}

type Params struct {
	fx.In
	I18n *ut.UniversalTranslator `name:"universalTranslator"`
}

func NewI18nService(params Params) II18nService {
	return &I18nService{
		i18n: params.I18n,
	}
}

func StringToPluralRule(s string) locales.PluralRule {
	switch s {
	case "Unknown":
		return locales.PluralRuleUnknown
	case "Zero":
		return locales.PluralRuleZero
	case "One":
		return locales.PluralRuleOne
	case "Two":
		return locales.PluralRuleTwo
	case "Few":
		return locales.PluralRuleFew
	case "Many":
		return locales.PluralRuleMany
	case "Other":
		return locales.PluralRuleOther
	default:
		return locales.PluralRuleUnknown
	}
}
