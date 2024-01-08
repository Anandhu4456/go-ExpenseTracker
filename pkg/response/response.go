package response

type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
	Error      error
}

func (r *Response) ClientResponse(statusCode int, message string, data interface{}, err error) *Response {
	return &Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}
}
