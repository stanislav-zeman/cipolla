package dto

import (
	"errors"
)

var ErrUnknownLogger = errors.New("unknown logger")

type Logger struct {
	Struct  string
	Package string
}

func ParseLogger(name string) (logger Logger, err error) {
	switch name {
	case "slog":
		return Logger{
			Struct:  "*slog.Logger",
			Package: "slog",
		}, nil

	case "zap":
		return Logger{
			Struct:  "*zap.Logger",
			Package: "go.uber.org/zap",
		}, nil

	case "zerolog":
		return Logger{
			Struct:  "*zerolog.Logger",
			Package: "github.com/rs/zerolog",
		}, nil

	default:
		return Logger{}, ErrUnknownLogger
	}
}
