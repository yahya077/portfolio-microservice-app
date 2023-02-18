package service

import (
	"context"
	"github.com/yahya077/portfolio-microservice-app/web-bff/contract"
)

const (
	KeySchema         = "schema"
	KeyCode           = "code"
	CodeSuccess       = 200
	CodeInternalError = 500
)

type ServiceContext struct {
	ctx context.Context
}

func (c *ServiceContext) GetSchema() interface{} {
	return c.ctx.Value(KeySchema)
}

func (c *ServiceContext) SetSchema(schema interface{}) contract.IServiceContext {
	c.ctx = context.WithValue(c.ctx, KeySchema, schema)
	return c
}

func (c *ServiceContext) GetCode() int {
	if v := c.ctx.Value(KeyCode); v != nil {
		return v.(int)
	}
	return 0
}

func (c *ServiceContext) SetCode(code int) contract.IServiceContext {
	c.ctx = context.WithValue(c.ctx, KeyCode, code)
	return c
}

func (c *ServiceContext) HasError() bool {
	return c.GetCode() != CodeSuccess
}

func (c *ServiceContext) GetError() error {
	return c.GetSchema().(error)
}

type Option func(f contract.IServiceContext)

func Context(opts ...Option) contract.IServiceContext {
	r := &ServiceContext{context.Background()}

	for _, opt := range opts {
		opt(r)
	}

	return r
}
