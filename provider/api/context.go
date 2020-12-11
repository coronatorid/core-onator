package api

import (
	"context"

	"github.com/labstack/echo/v4"
)

// Context of API
type Context struct {
	echo.Context
}

// FabricateContext for api handler
func FabricateContext(ctx echo.Context) *Context {
	return &Context{ctx}
}

// Ctx returning request context
func (c *Context) Ctx() context.Context {
	return c.Request().Context()
}
