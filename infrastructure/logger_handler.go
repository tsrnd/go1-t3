package infrastructure

import (
	"os"

	"github.com/sirupsen/logrus"
)

const (
	// OutputStdout check config(logger.output).
	OutputStdout = "stdout"
	// OutputFile check config(logger.output).
	OutputFile   = "file"

	// FormatText check config(logger.format).
	FormatText = "text"
	// FormatJSON check config(logger.format).
	FormatJSON = "json"
)

// LoggerHandler struct.
type LoggerHandler struct {
	Log     *logrus.Logger
	Logfile *os.File
}

// NewLoggerHandler returns new LoggerHandler.
// repository: https://github.com/sirupsen/logrus
func NewLoggerHandler() *LoggerHandler {
	var err error
	var file *os.File

	// get config.
	output := GetConfigString("logger.output")
	level := GetConfigString("logger.level")
	format := GetConfigString("logger.format")

	// new logrus.
	log := logrus.New()

	// set output.
	switch output {
	case OutputStdout: // output: stdout
		log.Out = os.Stdout
	case OutputFile: // output: file
		logfile := GetConfigString("logger.file")
		file, err = os.OpenFile(logfile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		log.Out = file
	}

	// set level.
	log.Level, err = logrus.ParseLevel(level)
	if err != nil {
		panic(err)
	}

	// set formatter.
	switch format {
	case FormatText:
		log.Formatter = &logrus.TextFormatter{}
	case FormatJSON:
		log.Formatter = &logrus.JSONFormatter{}
	}
	return &LoggerHandler{Log: log, Logfile: file}
}
