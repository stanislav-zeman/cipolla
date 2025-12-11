package response

import (
    "github.com/stanislav-zeman/cipolla/example/internal/content/application/query"

)

type GetProgramResponse struct{}

func NewGetProgramResponse(r query.ProgramQueryResult) (res GetProgramResponse, err error) {
    return GetProgramResponse{}, nil
}
