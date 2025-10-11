package response

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/application/command"
)

type PostStreamResponse struct{}

func NewPostStreamResponse(r command.StreamCreateCommandResult) (res PostStreamResponse, err error) {
    return PostStreamResponse{}, nil
}
