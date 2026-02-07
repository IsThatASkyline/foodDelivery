package common

// Response ответ со списком элементов
type Response struct {
	Status int         `json:"status"`
	Error  string      `json:"error,omitempty"`
	Data   interface{} `json:"data,omitempty"`
}

// DefaultResponse типовой ответ
type DefaultResponse struct {
	Status  int    `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ListData содержит список элементов ответа
type ListData struct {
	Items interface{} `json:"items"`
	Range ListRange   `json:"range"`
}

// ListRange описывает пейджер в ответе
type ListRange struct {
	Count  int64 `json:"count"`
	Limit  int64 `json:"limit"`
	Offset int64 `json:"offset"`
}
