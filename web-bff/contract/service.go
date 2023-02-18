package contract

import "github.com/gofiber/fiber/v2"

type IService interface {
	HandleGet(ctx *fiber.Ctx) IServiceContext
}

type IServiceContext interface {
	GetSchema() interface{}
	SetSchema(schema interface{}) IServiceContext
	GetCode() int
	SetCode(code int) IServiceContext
	HasError() bool
	GetError() error
}
