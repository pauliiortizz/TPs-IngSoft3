package dao

import (
	"time"
)

type Inscripciones struct {
	IdUsuario        int       `gorm:"column:Id_usuario;primaryKey;not null"`
	IdCurso          int       `gorm:"column:Id_curso;primaryKey;not null"`
	FechaInscripcion time.Time `gorm:"column:fecha_inscripcion;not null"`
	Comentario       string    `gorm:"column:comentario"`

	// Definir las relaciones de clave foránea con las tablas User y Curso
	User  User  `gorm:"foreignKey:IdUsuario"`
	Curso Curso `gorm:"foreignKey:IdCurso"`
}
