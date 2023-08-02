package logs_cm

import (
	"fmt"
	"time"

	"github.com/CarlosMore29/logs_cm/api"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

func NewLogger(logConfig LoggerConfig) *Logger {

	logger := &Logger{
		input: api.Input{
			Fields: api.Fields{
				Destino:   logConfig.Destination,
				Direccion: logConfig.Host,
			},
			Measurement: logConfig.Measurement,
			Tags: api.Tags{
				TipoServicio: logConfig.ModuleName,
			},
		},
		logger:       getInstanceLogger(),
		apiQLUrl:     logConfig.ApiQlUrl,
		apiQLSecret:  logConfig.ApiQlSecret,
		sendToServer: logConfig.SendToServer,
	}

	return logger
}

type LoggerConfig struct {
	ApiQlUrl     string
	ApiQlSecret  string
	Destination  string
	Host         string
	ModuleName   string
	Measurement  string
	SendToServer bool
}

type Logger struct {
	input        api.Input
	logger       *logrus.Logger
	apiQLUrl     string
	apiQLSecret  string
	sendToServer bool
}

func (l *Logger) Info(msg string) error {
	device := "System"
	extra := ""
	errLog := l.sendLogsExec(msg, "Info", device, extra)
	l.logger.Info(msg)
	return errLog
}
func (l *Logger) InfoDevice(msg, device string) error {
	extra := ""
	errLog := l.sendLogsExec(msg, "Info", device, extra)
	l.logger.Info(msg)
	return errLog
}
func (l *Logger) InfoExtra(msg, extra string) error {
	device := "System"
	errLog := l.sendLogsExec(msg, "Info", device, extra)
	l.logger.Info(msg)
	return errLog
}
func (l *Logger) InfoDeviceExtra(msg, device, extra string) error {
	errLog := l.sendLogsExec(msg, "Info", device, extra)
	l.logger.Info(msg)
	return errLog
}

func (l *Logger) Error(msg string) error {
	device := "System"
	extra := ""
	errLog := l.sendLogsExec(msg, "Error", device, extra)
	l.logger.Error(msg)
	return errLog
}
func (l *Logger) ErrorDevice(msg, device string) error {
	extra := ""
	errLog := l.sendLogsExec(msg, "Error", device, extra)
	l.logger.Error(msg)
	return errLog
}
func (l *Logger) ErrorExtra(msg, extra string) error {
	device := "System"
	errLog := l.sendLogsExec(msg, "Error", device, extra)
	l.logger.Error(msg)
	return errLog
}
func (l *Logger) ErrorDeviceExtra(msg, device, extra string) error {
	errLog := l.sendLogsExec(msg, "Error", device, extra)
	l.logger.Error(msg)
	return errLog
}

func (l *Logger) Debug(msg string) error {
	device := "System"
	extra := ""
	errLog := l.sendLogsExec(msg, "Debug", device, extra)
	l.logger.Debug(msg)
	return errLog
}
func (l *Logger) DebugDevice(msg, device string) error {
	extra := ""
	errLog := l.sendLogsExec(msg, "Debug", device, extra)
	l.logger.Debug(msg)
	return errLog
}
func (l *Logger) DebugExtra(msg, extra string) error {
	device := "System"
	errLog := l.sendLogsExec(msg, "Debug", device, extra)
	l.logger.Debug(msg)
	return errLog
}
func (l *Logger) DebugDeviceExtra(msg, device, extra string) error {
	errLog := l.sendLogsExec(msg, "Debug", device, extra)
	l.logger.Debug(msg)
	return errLog
}

func (l *Logger) Warn(msg string) error {
	device := "System"
	extra := ""
	errLog := l.sendLogsExec(msg, "Warn", device, extra)
	l.logger.Warn(msg)
	return errLog
}
func (l *Logger) WarnDevice(msg, device string) error {
	extra := ""
	errLog := l.sendLogsExec(msg, "Warn", device, extra)
	l.logger.Warn(msg)
	return errLog
}
func (l *Logger) WarnExtra(msg, extra string) error {
	device := "System"
	errLog := l.sendLogsExec(msg, "Warn", device, extra)
	l.logger.Warn(msg)
	return errLog
}
func (l *Logger) WarnDeviceExtra(msg, device, extra string) error {
	errLog := l.sendLogsExec(msg, "Warn", device, extra)
	l.logger.Warn(msg)
	return errLog
}

func (l *Logger) Fatal(msg string) {
	device := "System"
	extra := ""
	l.sendLogsExec(msg, "Fatal", device, extra)
	l.logger.Fatal(msg)
}
func (l *Logger) FatalDevice(msg, device string) {
	extra := ""
	l.sendLogsExec(msg, "Fatal", device, extra)
	l.logger.Fatal(msg)
}
func (l *Logger) FatalExtra(msg, extra string) {
	device := "System"
	l.sendLogsExec(msg, "Fatal", device, extra)
	l.logger.Fatal(msg)
}
func (l *Logger) FatalDeviceExtra(msg, device, extra string) {
	l.sendLogsExec(msg, "Fatal", device, extra)
	l.logger.Fatal(msg)
}

func (l *Logger) Panic(msg string) {
	device := "System"
	extra := ""
	l.sendLogsExec(msg, "Panic", device, extra)
	l.logger.Panic(msg)
}
func (l *Logger) PanicDevice(msg, device string) {
	extra := ""
	l.sendLogsExec(msg, "Panic", device, extra)
	l.logger.Panic(msg)
}
func (l *Logger) PanicExtra(msg, extra string) {
	device := "System"
	l.sendLogsExec(msg, "Panic", device, extra)
	l.logger.Panic(msg)
}
func (l *Logger) PanicDeviceExtra(msg, device, extra string) {
	l.sendLogsExec(msg, "Panic", device, extra)
	l.logger.Panic(msg)
}

func (l *Logger) Trace(msg string) error {
	device := "System"
	extra := ""
	errLog := l.sendLogsExec(msg, "Trace", device, extra)
	l.logger.Trace(msg)
	return errLog
}
func (l *Logger) TraceDevice(msg, device string) error {
	extra := ""
	errLog := l.sendLogsExec(msg, "Trace", device, extra)
	l.logger.Trace(msg)
	return errLog
}
func (l *Logger) TraceExtra(msg, extra string) error {
	device := "System"
	errLog := l.sendLogsExec(msg, "Trace", device, extra)
	l.logger.Trace(msg)
	return errLog
}
func (l *Logger) TraceDeviceExtra(msg, device, extra string) error {
	errLog := l.sendLogsExec(msg, "Trace", device, extra)
	l.logger.Trace(msg)
	return errLog
}

func (l *Logger) sendLogsExec(msg, typeServicio, device, extra string) error {

	var errSave error = nil

	if l.sendToServer {
		l.input.Fields.Data = msg
		l.input.Fields.Dispositivo = device
		l.input.Fields.FechaRegistro = time.Now()
		l.input.Fields.Extra = extra

		l.input.Tags.TipoLog = typeServicio

		errSave = api.SaveLogs(l.apiQLUrl, l.apiQLSecret, l.input)
		if errSave != nil {
			l.logger.Error(fmt.Sprintf("Error al guardar el siguiente log: %s | error: %s", l.input.ToString(), errSave.Error()))
		}
	}

	return errSave
}

func getInstanceLogger() *logrus.Logger {

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
