package response

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/application/query"

)

type GetConfigurationResponse struct{}

func NewGetConfigurationResponse(r query.ConfigurationQueryResult) (res GetConfigurationResponse, err error) {
    return GetConfigurationResponse{}, nil
}
