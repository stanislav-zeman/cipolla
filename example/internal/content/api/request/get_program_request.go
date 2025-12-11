package request

import (
    "github.com/stanislav-zeman/cipolla/example/internal/content/application/query"

    "github.com/labstack/echo/v4"
)

type GetProgramRequest struct{}

func ParseGetProgramRequest(c *echo.Context) (r GetProgramRequest, err error) {
    return GetProgramRequest{}, nil
}

func (r *GetProgramRequest) ToQuery() (d query.ProgramQuery, err error) {
    return
}
