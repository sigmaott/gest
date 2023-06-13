package loader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
)

type II18nLoader interface {
	LoadData() map[string]ListTranslation
}
type ListTranslation []Translation
type I18nJsonLoader struct {
	path string
	data map[string]ListTranslation
}

func (i *I18nJsonLoader) LoadData() map[string]ListTranslation {

	return i.data
}

type Params struct {
	Path string
}

type Translation struct {
	Key      string `json:"key"`
	Trans    string `json:"trans"`
	Type     string `json:"type"`
	Rule     string `json:"rule"`
	Override bool   `json:"override"`
}

func (i *I18nJsonLoader) loadTranslations(path string) error {
	translations := make(map[string]ListTranslation)

	files, err := filepath.Glob(filepath.Join(i.path, "*.json"))
	if err != nil {
		return err
	}

	for _, file := range files {
		log.Printf("json: %s", file)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		var trans []Translation
		if err := json.Unmarshal(data, &trans); err != nil {
			return err
		}
		log.Printf("json: %s", data)
		fileExt := filepath.Ext(file)
		fileName := filepath.Base(file)
		translations[fileName[0:len(fileName)-len(fileExt)]] = trans

	}
	i.data = translations
	return nil
}
func NewI18nJsonLoader(param Params) II18nLoader {
	i := &I18nJsonLoader{
		path: param.Path,
		data: map[string]ListTranslation{},
	}
	err := i.loadTranslations(i.path)
	if err != nil {
		fmt.Println("i18n err")
		fmt.Println(err)
	}
	return i
}
