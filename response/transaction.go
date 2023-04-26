package response

import "net/http"

type Response struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

func NewSuccessResponse(data interface{}) Response {
	return Response{
		Code:   http.StatusOK,
		Status: http.StatusText(http.StatusOK),
		Data:   data,
	}
}

func NewErrorResponse(code int, err error) Response {
	return Response{
		Code:   code,
		Status: err.Error(),
		Data:   nil,
	}
}
