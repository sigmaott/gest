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
	GetListBaseQuota(quotaGroups []string) ([]*model.GroupQuota, error)
	GetBaseQuota(name string) (*model.GroupQuota, error)
}

type baseQuotaRepository struct {
	groupQuota map[string]model.GroupQuota
}

func (b *baseQuotaRepository) GetListBaseQuota(quotaGroupNames []string) ([]*model.GroupQuota, error) {
	s := lo.MapToSlice(b.groupQuota, func(k string, v model.GroupQuota) *model.GroupQuota {
		return &v
	})
	if len(quotaGroupNames) == 0 {
		return s, nil
	}
	return lo.Filter(s, func(item *model.GroupQuota, index int) bool {
		return lo.Contains(quotaGroupNames, item.Name)
	}), nil

}

func (b *baseQuotaRepository) GetBaseQuota(name string) (*model.GroupQuota, error) {

	quotaGroup, ok := b.groupQuota[name]
	if !ok {
		return nil, errors.New("group quota not found")
	}
	return &quotaGroup, nil
}

func NewBaseQuotaRepository() BaseQuotaRepository {

	data, err := ioutil.ReadFile(config2.GetConfiguration().QuotaPath)
	if err != nil {
		return nil
	}
	configQuota := struct {
		GroupQuota []model.GroupQuota `yaml:"group_quota"`
	}{}
	err = yaml.Unmarshal(data, &configQuota)
	if err != nil {
		return nil
	}

	return &baseQuotaRepository{
		groupQuota: lo.Associate(configQuota.GroupQuota, func(f model.GroupQuota) (string, model.GroupQuota) {
			return f.Name, f
		}),
	}

}
