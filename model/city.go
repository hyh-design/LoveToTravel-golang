package model

type City struct {
	CityId       string `form:"cityId" json:"cityId"`
	CityName     string `form:"cityName" json:"cityName"`
	CityEname    string `form:"cityEname" json:"cityEname"`
	Lng          string `form:"lng" json:"lng"`
	Lat          string `form:"lat" json:"lat"`
	Url          string `form:"url" json:"url"`
	Introduction string `form:"introduction" json:"introduction"`
}

func (v City) TableName() string {
	return "city"
}
