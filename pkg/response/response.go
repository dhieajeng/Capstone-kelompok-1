package response

type Response struct {
	Code    int         `json:"code"`
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(code int, isSuccess bool, message string, data interface{}) Response {
	return Response{
		Code:    code,
		Success: isSuccess,
		Message: message,
		Data:    data,
	}
}

func Error(code int, isSuccess bool, message interface{}) Response {
	return Response{
		Code:    code,
		Success: isSuccess,
		Message: message,
		Data:    nil,
	}
}
