package helper

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type EmptyObject struct{}

func BuildReponse(status string, message string, data interface{}) Response {
	res := Response{
		Code:    status,
		Message: message,
		Data:    data,
	}

	return res
}
