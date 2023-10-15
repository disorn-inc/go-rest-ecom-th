package entities

import "github.com/disorn-inc/go-rest-ecom-th/pkg/router"

type IResponse interface {
	Succeed(code int, data any) IResponse
	Error(code int, traceId string, message string) IResponse
	Response()
}

type Response struct {
	StatusCode int            `json:"statusCode"`
	ErrorRes   *ErrorResponse `json:"errorResponse"`
	Data       any            `json:"data"`
	IsError    bool           `json:"isError"`
	Context    router.Context
}

type ErrorResponse struct {
	TraceId string `json:"traceId"`
	Message string `json:"message"`
}

func NewResponseV2(code int, data any, traceId string, message string) *Response {
	var errRes *ErrorResponse
	if traceId == "" && message == "" {
		errRes = nil
	} else {
		errRes = &ErrorResponse{
			TraceId: traceId,
			Message: message,
		}
	}
	return &Response{
		StatusCode: code,
		ErrorRes:   errRes,
		Data:       data,
		IsError:    errRes != nil,
	}
}

func NewResponse(ctx router.Context) IResponse {
	return &Response{
		Context: ctx,
	}
}

func (r *Response) Succeed(code int, data any) IResponse {
	r.StatusCode = code
	r.Data = data
	r.IsError = false
	return r
}

func (r *Response) Error(code int, traceId string, message string) IResponse {
	r.StatusCode = code
	r.ErrorRes = &ErrorResponse{
		TraceId: traceId,
		Message: message,
	}
	r.IsError = true
	return r
}

func (r *Response) Response() {
	if r.IsError {
		r.Context.JSON(r.StatusCode, r.ErrorRes)
	}
	r.Context.JSON(r.StatusCode, r.Data)
}