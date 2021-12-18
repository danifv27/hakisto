// Package logrus implements a adidas.com/cmo/services/k8ssecretum/log.Logger using Logrus as a backend
// You can use this by creating a logrus logger and calling `FromLogrus(entry)`.
// If you want this to be the default logger, set `log.L` to the value returned by `FromLogrus`
package logrus

import (
	"io"

	"github.com/danifv27/hakisto"
	"github.com/sirupsen/logrus"
)

// Adapter implements the `log.Logger` interface for logrus
type Adapter struct {
	//by default the name of an anonymous field is the name of its type
	*logrus.Entry
}

// NewLogrusLogger creates and configure a new logger
func NewLogrusLogger(out io.Writer, verbosity string, json bool) (*logrus.Logger, error) {

	l := logrus.New()
	if json {
		l.Formatter = new(logrus.JSONFormatter)
	}
	l.SetOutput(out)
	lvl, err := logrus.ParseLevel(verbosity)
	if err != nil {
		return nil, err
	}
	l.SetLevel(lvl)

	return l, nil
}

// FromLogrus creates a new `log.Logger` from the provided entry
func FromLogrus(entry *logrus.Entry) hakisto.Logger {

	return &Adapter{entry}
}

// WithField adds a field to the log entry.
func (l *Adapter) WithField(key string, val interface{}) hakisto.Logger {
	return FromLogrus(l.Entry.WithField(key, val))
}

// WithFields adds multiple fields to a log entry.
func (l *Adapter) WithFields(f hakisto.Fields) hakisto.Logger {
	return FromLogrus(l.Entry.WithFields(logrus.Fields(f)))
}

// WithError adds an error to the log entry
func (l *Adapter) WithError(err error) hakisto.Logger {
	return FromLogrus(l.Entry.WithError(err))
}
