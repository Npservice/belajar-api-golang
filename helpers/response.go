package helpers

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type ResponseError struct {
	Meta  Meta        `json:"meta"`
	Error interface{} `json:"error"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ApiResponse(message string, code int, status string, Data interface{}) Response {
	meta := Meta{
		Message: message,
		Code:    code,
		Status:  status,
	}
	response := Response{
		Meta: meta,
		Data: Data,
	}
	return response

}
