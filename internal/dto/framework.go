package dto

import "errors"

var ErrUnknownFramework = errors.New("unknown framework")

type Framework struct {
	Package       string
	ContextName   string
	ContextStruct string
}

func ParseFramework(name string) (framework Framework, err error) {
	switch name {
	case "gin":
		return Framework{
			ContextName:   "c",
			ContextStruct: "*gin.Context",
			Package:       "github.com/gin-gonic/gin",
		}, nil

	case "echo":
		return Framework{
			ContextName:   "c",
			ContextStruct: "*echo.Context",
			Package:       "github.com/labstack/echo/v4",
		}, nil

	case "http":
		return Framework{
			ContextName:   "r",
			ContextStruct: "*http.Request",
			Package:       "net/http",
		}, nil

	default:
		return Framework{}, ErrUnknownFramework
	}
}
