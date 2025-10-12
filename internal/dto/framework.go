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
		framework = Framework{
			ContextName:   "c",
			ContextStruct: "*gin.Context",
			Package:       "slog",
		}

		return framework, err

	case "echo":
		framework = Framework{
			ContextName:   "c",
			ContextStruct: "*echo.Context",
			Package:       "github.com/labstack/echo/v4",
		}

		return framework, err

	case "http":
		framework = Framework{
			ContextName:   "r",
			ContextStruct: "*http.Request",
			Package:       "net/http",
		}

		return framework, err

	default:
		err = ErrUnknownFramework
		return framework, err
	}
}
