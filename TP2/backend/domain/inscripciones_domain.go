package domain

import "time"

type Inscripcion struct {
	IdUsuario        int       `json:"id_usuario"`
	Usuario          User      `json:"usuario"`
	IdCurso          int       `json:"id_curso"`
	Curso            Curso     `json:"curso"`
	FechaInscripcion time.Time `json:"fecha_inscripcion"`
	Comentario       string    `json:"comentario"`
}
