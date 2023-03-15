package model

type Province struct {
	ProvinceId   string `form:"provinceId" json:"provinceId"`
	ProvinceName string `form:"provinceName" json:"provinceName"`
	Url          string `form:"url" json:"url"`
	Introduction string `form:"introduction" json:"introduction"`
}

func (v Province) TableName() string {
	return "province"
}
