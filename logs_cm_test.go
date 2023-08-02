package logs_cm_test

import (
	"os"
	"testing"

	"github.com/CarlosMore29/env_cm"
	"github.com/CarlosMore29/logs_cm"
)

func envGLobals() {
	env_cm.GetEnvFile()
}

func TestGeneral(t *testing.T) {
	envGLobals()
	dispositivo := "DEV001"
	loggerconfig := logs_cm.LoggerConfig{
		ApiQlUrl:     os.Getenv("APIQL_URL"),
		ApiQlSecret:  os.Getenv("APIQL_SECRET"),
		Destination:  "test",
		Host:         "localhost",
		ModuleName:   "Modulo Test",
		Measurement:  "system_test",
		SendToServer: true,
	}
	logger := logs_cm.NewLogger(loggerconfig)
	errlogDevice := logger.InfoDevice("Esto si es un mensaje Tipo INFO de Log Device", dispositivo)
	if errlogDevice != nil {
		t.Errorf(errlogDevice.Error())
	}
	errlog := logger.Info("Esto si es un mensaje Tipo INFO de Log System")
	if errlog != nil {
		t.Errorf(errlog.Error())
	}

	errLogErr := logger.Error("Esto si es un mensaje Tipo ERROR de Log System")
	if errLogErr != nil {
		t.Errorf(errLogErr.Error())
	}
	errLogErrDispo := logger.ErrorDevice("Esto si es un mensaje Tipo ERROR de Log System", dispositivo)
	if errLogErrDispo != nil {
		t.Errorf(errLogErrDispo.Error())
	}

}
