package models

type StatisticResponse struct {
	Likes     int `json:"likes"`
	ViewCount int `json:"viewCount"`
	Contacts  int `json:"contacts"`
}

type ErrorResponse struct {
	Status string `json:"status"`
	Result struct {
		Message  string                 `json:"message"`
		Messages map[string]interface{} `json:"messages"`
	} `json:"result"`
}

type SimpleErrorResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}
