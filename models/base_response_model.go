package models

type APIResponse struct {
	Success bool        `json:"success"`
	Status  int         `json:"status"`
	Code    ErrorCode   `json:"code,omitempty"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func Success(status int, message string, data interface{}) APIResponse {
	return APIResponse{
		Success: true,
		Status:  status,
		Message: message,
		Data:    data,
	}
}

func Error(status int, code ErrorCode, message string) APIResponse {
	return APIResponse{
		Success: false,
		Status:  status,
		Code:    code,
		Message: message,
		Data:    nil,
	}
}
