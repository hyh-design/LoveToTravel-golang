package model

import "gorm.io/gorm"

type Scenery struct {
	gorm.Model
	Id           string `form:"id" json:"id"`
	Name         string `form:"name" json:"name"`
	Introduction string `form:"introduction" json:"introduction"`
	Score        string `form:"score" json:"score"`
	Ticket       string `form:"ticket" json:"ticket"`
	Opening      string `form:"opening" json:"opening"`
	Lng          string `form:"lng" json:"lng"`
	Lat          string `form:"lat" json:"lat"`
	Level        string `form:"level" json:"level"`
	Address      string `form:"address" json:"address"`
	Season       string `form:"season" json:"season"`
	Tips         string `form:"tips" json:"tips"`
	CityName     string `form:"cityName" json:"cityName"`
	CityId       string `form:"cityId" json:"cityId"`
	Url          string `form:"url" json:"url"`
	Tele         string `form:"tele" json:"tele"`
	Cluster      string `form:"cluster" json:"cluster"`
}

func (v Scenery) TableName() string {
	return "scenery"
}
