package main

import (
	"context"
	"os"

	log "github.com/danifv27/hakisto"
	logruslogger "github.com/danifv27/hakisto/logrus"
	"github.com/sirupsen/logrus"
)

func main() {
	var ctx context.Context
	var err error
	var l *logrus.Logger

	ctx, _ = context.WithCancel(context.Background())

	if l, err = logruslogger.NewLogrusLogger(os.Stdout, "Debug", false); err != nil {
		l = logrus.StandardLogger()
	}
	log.L = logruslogger.FromLogrus(logrus.NewEntry(l))

	log.G(ctx).WithFields(log.Fields{
		"field1": "value1",
	}).Debug("debug test message")
}
