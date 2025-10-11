package request

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/application/command"
    "github.com/labstack/echo/v4"
)

type PostStreamRequest struct{}

func ParsePostStreamRequest(c *echo.Context) (r PostStreamRequest, err error) {
    return PostStreamRequest{}, nil
}

func (r *PostStreamRequest) ToCommand() (d command.StreamCreateCommand, err error) {
    return
}
