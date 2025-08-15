package domain

import (
	"encoding/json"
	"time"
)

type Curso struct {
	IdCurso     int       `json:"id_curso"`
	Titulo      string    `json:"Titulo"`
	FechaInicio time.Time `json:"FechaInicio"`
	Categoria   string    `json:"Categoria"`
	Archivo     string    `json:"Archivo"`
	Descripcion string    `json:"Descripcion"`
	Usuarios    []User    `json:"-"`
}

type SearchResponse struct {
	Results []Curso `json:"results"`
}

type SubscribeRequest struct {
	IdUsuario int `json:"id_usuario"`
	IdCurso   int `json:"id_curso"`
}

// UnmarshalJSON es necesario para deserializar FechaInicio en formato "2006-01-02"
func (c *Curso) UnmarshalJSON(data []byte) error {
	type Alias Curso
	aux := &struct {
		FechaInicio string `json:"FechaInicio"`
		*Alias
	}{
		Alias: (*Alias)(c),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	// Parsear FechaInicio
	date, err := time.Parse("2006-01-02", aux.FechaInicio)
	if err != nil {
		return err
	}
	c.FechaInicio = date

	return nil
}

// MarshalJSON es necesario para serializar FechaInicio en formato "2006-01-02"
func (c *Curso) MarshalJSON() ([]byte, error) {
	type Alias Curso
	return json.Marshal(&struct {
		FechaInicio string `json:"FechaInicio"`
		*Alias
	}{
		FechaInicio: c.FechaInicio.Format("2006-01-02"),
		Alias:       (*Alias)(c),
	})
}
