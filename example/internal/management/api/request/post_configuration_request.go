package request

import (
    "github.com/stanislav-zeman/cipolla/example/internal/management/application/command"
    "github.com/labstack/echo/v4"
)

type PostConfigurationRequest struct{}

func ParsePostConfigurationRequest(c *echo.Context) (r PostConfigurationRequest, err error) {
    return PostConfigurationRequest{}, nil
}

func (r *PostConfigurationRequest) ToCommand() (d command.ConfigurationCreateCommand, err error) {
    return
}
