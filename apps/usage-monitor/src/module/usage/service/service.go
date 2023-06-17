package service

import "usage-monnitor/src/module/usage/repository"

type ISSAIUsageMonitorService interface {
	CheckQuota(appId string, resource string) (bool, error)
}
type ssaiUsageMonitorService struct {
	quotaRepository repository.IQuotaRepository
	usageRepository repository.IUsageRepository
}

func (s *ssaiUsageMonitorService) CheckQuota(appId string, resource string) (bool, error) {
	usage, err := s.usageRepository.GetUsageByMetric(appId, resource)
	if err != nil {
		return false, err
	}
	quota, err := s.quotaRepository.GetQuotaResource(appId, resource)
	if err != nil {
		return false, err
	}
	return usage.Usage <= quota.Quota.Hard, nil

}

func NewSSAIUsageMonitorService(
	quotaRepository repository.IQuotaRepository,
	usageRepository repository.IUsageRepository) ISSAIUsageMonitorService {
	return &ssaiUsageMonitorService{
		quotaRepository: quotaRepository,
		usageRepository: usageRepository,
	}
}
