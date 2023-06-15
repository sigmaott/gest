package repository

import (
	"errors"
	"github.com/samber/lo"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	config2 "quota/config"
	"quota/src/module/quota/model"
)

type name struct {
	GroupQuota map[string]model.GroupQuota
}
type BaseQuotaRepository interface {
	GetBaseQuota(name string) (*model.GroupQuota, error)
}

type baseQuotaRepository struct {
	groupQuota map[string]model.GroupQuota
}

func (b *baseQuotaRepository) GetBaseQuota(name string) (*model.GroupQuota, error) {
	if len(b.groupQuota) == 0 {
		data, err := ioutil.ReadFile(config2.GetConfiguration().QuotaPath)
		if err != nil {
			return nil, err
		}
		configQuota := struct {
			GroupQuota []model.GroupQuota `yaml:"group_quota"`
		}{}
		err = yaml.Unmarshal(data, &configQuota)
		if err != nil {
			return nil, err
		}
		b.groupQuota = lo.Associate(configQuota.GroupQuota, func(f model.GroupQuota) (string, model.GroupQuota) {
			return f.Name, f
		})
	}
	quotaGroup, ok := b.groupQuota[name]
	if !ok {
		return nil, errors.New("group quota not found")
	}
	return &quotaGroup, nil
}

func NewBaseQuotaRepository() BaseQuotaRepository {
	return &baseQuotaRepository{}

}
