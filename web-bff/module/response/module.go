package response

import (
	"github.com/yahya077/portfolio-microservice-app/web-bff/contract"
)

const (
	MessageSuccess = "Successfully Finished"
	CodeSuccess    = 200
	CodeError      = 500
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func (r *Response) GetResponse() interface{} {
	return r
}

func (r *Response) GetCode() int {
	return r.Code
}

func (r *Response) SetCode(code int) contract.IResponse {
	r.Code = code
	return r
}

func (r *Response) SetMessage(message string) contract.IResponse {
	r.Message = message
	return r
}

func (r *Response) SetData(data interface{}) contract.IResponse {
	r.Data = data
	return r
}

type Option func(f contract.IResponse)

func NewResponse(opts ...Option) *Response {
	r := &Response{}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

func WithError(e error) Option {
	return func(f contract.IResponse) {
		f = f.SetMessage(e.Error())
	}
}

func WithSuccess() Option {
	return func(f contract.IResponse) {
		f = f.SetMessage(MessageSuccess).SetCode(200)
	}
}

func WithStatusError() Option {
	return func(f contract.IResponse) {
		f = f.SetCode(CodeError)
	}
}

func WithStatusSuccess() Option {
	return func(f contract.IResponse) {
		f = f.SetCode(CodeSuccess)
	}
}

func WithData(data interface{}) Option {
	return func(f contract.IResponse) {
		f = f.SetData(data)
	}
}
