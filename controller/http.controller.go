package controller

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Status  string      `json:"status"`
	Error   error       `json:"error"`
}

type MetaResponse struct {
	Data     []interface{} `json:"data"`
	Message  string        `json:"message"`
	Status   string        `json:"status"`
	Error    error         `json:"error"`
	Page     int           `json:"page"`
	PageSize int           `json:"pageSize"`
	Total    int           `json:"total"`
}
