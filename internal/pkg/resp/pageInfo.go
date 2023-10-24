package resp

type PageInfo struct {
	HasMore  bool        `json:"has_more"`
	Page     int32       `json:"page"`
	PageData interface{} `json:"page_data"`
}
