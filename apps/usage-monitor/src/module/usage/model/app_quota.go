package model

import "time"

type AppQuota struct {
	ID         string    `json:"id" bson:"id"`
	AppId      string    `json:"appId" bson:"appId"`
	GroupQuota string    `json:"groupQuota" bson:"groupQuota"`
	CreatedAt  time.Time `json:"createdAt" bson:"createdAt"`
	UpdateAt   time.Time `json:"createdAt" bson:"updateAt"`
}
