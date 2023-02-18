package contract

type IResponse interface {
	GetResponse() interface{}
	GetCode() int
	SetCode(code int) IResponse
	SetMessage(message string) IResponse
	SetData(data interface{}) IResponse
}
