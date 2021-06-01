package helper

type Presenter struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

func GenerateResponse(statusCode int, message string, data interface{}) *Presenter {
	return &Presenter{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
	}
}
