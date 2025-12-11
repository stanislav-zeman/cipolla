package request

import (
    "github.com/stanislav-zeman/cipolla/example/internal/management/application/command"
    "github.com/labstack/echo/v4"
)

type PatchStreamRequest struct{}

func ParsePatchStreamRequest(c *echo.Context) (r PatchStreamRequest, err error) {
    return PatchStreamRequest{}, nil
}

func (r *PatchStreamRequest) ToCommand() (d command.StreamUpdateCommand, err error) {
    return
}
