package request

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/application/query"

    "github.com/labstack/echo/v4"
)

type GetStreamRequest struct{}

func ParseGetStreamRequest(c *echo.Context) (r GetStreamRequest, err error) {
    return GetStreamRequest{}, nil
}

func (r *GetStreamRequest) ToQuery() (d query.StreamQuery, err error) {
    return
}
