package entity

import (
    "github.com/stanislav-zeman/cipolla/example/internal/management/domain/value"
)

type Stream struct {
    ID string
    Name string
    State value.State
}
