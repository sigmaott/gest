package model

type Quota struct {
	Resource string  `yaml:"resource"`
	Hard     float64 `yaml:"hard"`
}

type GroupQuota struct {
	Name   string           `yaml:"name"`
	Quotas map[string]Quota `yaml:"quotas"`
}
