package api

import (
	"fmt"
	"time"
)

type Input struct {
	Fields      Fields `json:"Fields"`
	Measurement string `json:"Measurement"`
	Tags        Tags   `json:"Tags"`
}

func (i *Input) ToString() string {
	return fmt.Sprintf("{fields: %s, Measurement: %s, tags: %s}", i.Fields.ToString(), i.Measurement, i.Tags.ToString())
}

type Fields struct {
	Data          string    `json:"data"`
	Destino       string    `json:"destino"`
	Direccion     string    `json:"direccion"`
	Dispositivo   string    `json:"dispositivo"`
	Extra         string    `json:"extra"`
	FechaRegistro time.Time `json:"fecha_registro"`
}

func (f *Fields) ToString() string {
	return fmt.Sprintf("{data: %s, destino: %s, direccion: %s, Dispositivo: %s, extra: %s, fechaRegistro: %s}",
		f.Data,
		f.Destino,
		f.Direccion,
		f.Dispositivo,
		f.Extra,
		f.FechaRegistro.Format(time.RFC3339),
	)
}

type Tags struct {
	TipoLog      string `json:"tipo_log"`
	TipoServicio string `json:"tipo_servicio"`
}

func (t *Tags) ToString() string {
	return fmt.Sprintf("{tipoLog: %s, tipoServicio: %s}", t.TipoLog, t.TipoServicio)
}
