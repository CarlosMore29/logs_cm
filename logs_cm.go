package logs_cm

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "./var/log/info_access.log",
		logrus.ErrorLevel: "./var/log/error_access.log",
		logrus.DebugLevel: "./var/log/debug_access.log",
		logrus.WarnLevel:  "./var/log/warning_access.log",
		logrus.FatalLevel: "./var/log/fatal_access.log",
		logrus.PanicLevel: "./var/log/panic_access.log",
		logrus.TraceLevel: "./var/log/trace_access.log",
	}

	Log := logrus.New()

	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))

	return Log
}
