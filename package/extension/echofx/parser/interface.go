package parser

import "github.com/labstack/echo/v4"

type IParser interface {
	Parser(next echo.HandlerFunc) echo.HandlerFunc
}
type Parser[T any] struct {
	name   string
	binder echo.Binder
}
