package helper

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func APIResponse(status string, data interface{}) Response {
	jsonResponse := Response{
		Status: status,
		Data:   data,
	}

	return jsonResponse
}

func TransAPIResponse(status, message string, data interface{}) Response {
	jsonTransResponse := Response{
		Status:  status,
		Message: message,
		Data:    data,
	}

	return jsonTransResponse
}
