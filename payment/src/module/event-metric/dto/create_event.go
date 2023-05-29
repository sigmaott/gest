package dto

type CreateEvent struct {
	TotalAdsInsert int64  `json:"total_ads_insert"`
	AppId          string `json:"app_id"`
}
