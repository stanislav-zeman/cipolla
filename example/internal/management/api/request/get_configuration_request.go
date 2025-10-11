package request

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/application/query"

    "github.com/labstack/echo/v4"
)

type GetConfigurationRequest struct{}

func ParseGetConfigurationRequest(c *echo.Context) (r GetConfigurationRequest, err error) {
    return GetConfigurationRequest{}, nil
}

func (r *GetConfigurationRequest) ToQuery() (d query.ConfigurationQuery, err error) {
    return
}
