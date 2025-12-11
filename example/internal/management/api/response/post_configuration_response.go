package response

import (
    "github.com/stanislav-zeman/cipolla/example/internal/management/application/command"
)

type PostConfigurationResponse struct{}

func NewPostConfigurationResponse(r command.ConfigurationCreateCommandResult) (res PostConfigurationResponse, err error) {
    return PostConfigurationResponse{}, nil
}
