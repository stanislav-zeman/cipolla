package response

import (
    "github.com/stanislav-zeman/cipolla/example/internal/management/application/command"
)

type PostStreamResponse struct{}

func NewPostStreamResponse(r command.StreamCreateCommandResult) (res PostStreamResponse, err error) {
    return PostStreamResponse{}, nil
}
