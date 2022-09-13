# README

Escribe archivos de logs de nivel:

* InfoLevel
* ErrorLevel
* DebugLevel
* WarnLevel
* FatalLevel
* PanicLevel
* TraceLevel

Los archivos se almacenan en la siguiente direecion apartir de la raiz

        ./var/log/

Apoyado por

* github.com/sirupsen/logrus
* github.com/rifflock/lfshook

## USO

Simple

```go
logger = logs_cm.NewLogger()
```
