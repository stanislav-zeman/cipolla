package dto

import (
	"net/http"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Request struct {
	Name               string
	Method             string
	ApplicationDTOName string

	Framework Framework `json:"-"`
	Import    Import    `json:"-"`
}

func NewRequest(name, method string, framework Framework, imp Import) Request {
	r := Request{
		Name:   name,
		Method: method,

		Framework: framework,
		Import:    imp,
	}

	switch strings.ToUpper(method) {
	case http.MethodGet:
		r.ApplicationDTOName = name + "Query"
	case http.MethodPost:
		r.ApplicationDTOName = name + "CreateCommand"
	case http.MethodPut:
		r.ApplicationDTOName = name + "ReplaceCommand"
	case http.MethodPatch:
		r.ApplicationDTOName = name + "UpdateCommand"
	case http.MethodDelete:
		r.ApplicationDTOName = name + "DeleteCommand"
	default:
		r.ApplicationDTOName = "Unknown"
	}

	r.Method = cases.Title(language.English).String(r.Method)

	return r
}
