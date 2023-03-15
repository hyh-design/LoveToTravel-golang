package vo

type Page struct {
	QueryStr string      `form:"queryStr" json:"queryStr"`
	PageNum  int         `form:"pageNum" json:"pageNum"`
	PageSize int         `form:"pageSize" json:"pageSize"`
	Total    int         `form:"total" json:"total"`
	PageSum  int         `form:"pageSum" json:"pageSum"`
	Records  interface{} `form:"records" json:"records"`
}
