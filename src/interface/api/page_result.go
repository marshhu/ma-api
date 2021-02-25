package api

type PageResult struct {
	Total int64       `json:"total"`
	Data  interface{} `json:"data"`
}
