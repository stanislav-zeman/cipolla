package dto

import (
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Response struct {
	Name   string
	Method string

	ApplicationDTOName string    `json:"-"`
	Framework          Framework `json:"-"`
	Import             Import    `json:"-"`
}

func NewResponse(name, method string, framework Framework, imp Import) Response {
	r := Response{
		Name:   name,
		Method: method,

		Framework: framework,
		Import:    imp,
	}

	switch strings.ToUpper(method) {
	case http.MethodGet:
		r.ApplicationDTOName = name + "QueryResult"
	case http.MethodPost:
		r.ApplicationDTOName = name + "CreateCommandResult"
	case http.MethodPut:
		r.ApplicationDTOName = name + "ReplaceCommandResult"
	case http.MethodPatch:
		r.ApplicationDTOName = name + "UpdateCommandResult"
	case http.MethodDelete:
		r.ApplicationDTOName = name + "DeleteCommandResult"
	default:
		r.ApplicationDTOName = "UnknownResult"
	}

	r.Method = cases.Title(language.English).String(r.Method)

	return r
}
