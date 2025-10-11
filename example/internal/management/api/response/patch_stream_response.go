package response

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/application/command"
)

type PatchStreamResponse struct{}

func NewPatchStreamResponse(r command.StreamUpdateCommandResult) (res PatchStreamResponse, err error) {
    return PatchStreamResponse{}, nil
}
