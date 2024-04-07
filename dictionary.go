package model

type AppDictionary struct {
	Id         int    `column:"id" json:"id"`
	Type       string `column:"type" json:"type"`
	Name       string `column:"name" json:"name"`
	Status     int    `column:"status" json:"status"`
	CreateTime string `column:"create_time" json:"create_time"`
}
