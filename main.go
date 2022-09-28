package main

import (
	"errors"
	pkgerr "github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"zerolog_stack/bootstrap"
)

type BizThing struct {
	logger zerolog.Logger
}

var inner func() error

// StandardLibStackTrace shows that if we use the standard lib error we don't get a stack trace
func StandardLibStackTrace() {

	// Get a standard lib error - note we don't see a stack trace
	withStdLibError()
	err := outer()
	log.Error().Stack().Err(err).Msg("Explicit stack call with stdlib error")

}

// ErrorPkgStackTrace shows if we use the errors pkg we do get a stack trace
func ErrorPkgStackTrace() {

	// Get a 3rd party lib error - note we do see a stack trace
	withPkgLibError()
	err := outer()
	log.Error().Stack().Err(err).Msg("Explicit stack call with 3rd party lib error")

}

func (b BizThing) HookDemo() {
	b.logger.Debug().Msg("Debug level message") // no line info in the log entry
	b.logger.Error().Msg("Error level message") // line info in the log entry
}

func withStdLibError() {
	inner = func() error {
		return errors.New("seems we have a standard lib error here")
	}
}

func withPkgLibError() {
	inner = func() error {
		return pkgerr.New("seems we have a 3rd party package error here")
	}
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	logger := bootstrap.SetupZeroLog()
	StandardLibStackTrace()
	ErrorPkgStackTrace()

	// Shows how we *could* use the hook functionality
	b := BizThing{logger: logger}
	b.HookDemo()
}
