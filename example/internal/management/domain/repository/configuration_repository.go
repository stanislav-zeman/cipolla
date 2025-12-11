package repository

import (
    "github.com/stanislav-zeman/cipolla/example/internal/management/domain/entity"
)

type ConfigurationRepository interface {
    GetAllConfigurations() ([]entity.Configuration, error)
    GetConfiguration(id string) (entity.Configuration, error)
    CreateConfiguration(entity.Configuration) (entity.Configuration, error)
    UpdateConfiguration(entity.Configuration) (entity.Configuration, error)
    DeleteConfiguration(id string) (entity.Configuration, error)
}

