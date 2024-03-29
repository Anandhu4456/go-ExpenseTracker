package response


type Response struct {
	StatusCode int
	Message    string
	Data       interface{}
	Error      interface{}
}

func ClientResponse(statusCode int, message string, data interface{}, err interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Data:       data,
		Error:      err,
	}
}
