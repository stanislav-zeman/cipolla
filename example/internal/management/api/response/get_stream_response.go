package response

import (
    "github.com/stanislav-zeman/gonion/example/internal/management/application/query"

)

type GetStreamResponse struct{}

func NewGetStreamResponse(r query.StreamQueryResult) (res GetStreamResponse, err error) {
    return GetStreamResponse{}, nil
}
