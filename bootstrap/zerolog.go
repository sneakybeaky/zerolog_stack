package bootstrap

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
)

type AddStackToErrors struct {
}

func (a AddStackToErrors) Run(e *zerolog.Event, level zerolog.Level, _ string) {
	if level >= zerolog.ErrorLevel {
		e.Stack()
		e.Caller(3)
	}
}

func SetupZeroLog() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	return log.Hook(AddStackToErrors{})
}
