package app

type BaseResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type ResponseData struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ErrorCodeResponse struct {
	StatusCode int
	Message    string `json:"message"`
}
