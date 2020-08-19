package context

type ListData struct {
	PageNo   int `c:"pageNo" json:"page_no"`     //当前页数
	PageSize int `c:"pageSize" json:"page_size"` //分页量
}
