package dto

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Config struct {
	Name string
}

func NewConfig(name string) Config {
	name = cases.Title(language.English).String(name)

	return Config{
		Name: name,
	}
}
