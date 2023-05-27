package model

import (
	"encoding/json"
	"log"
)

type SSAIAdsInsertProperties struct {
	TotalAdsInsert int64 `json:"total_ads_insert"`
}

func (obj *SSAIAdsInsertProperties) StructToMap() map[string]any {
	byteObj, err := json.Marshal(obj)
	if err != nil {
		log.Print(err)
		return nil
	}
	var result = map[string]any{}
	err = json.Unmarshal(byteObj, &result)
	if err != nil {
		log.Print(err)
		return nil
	}
	return result
}
