package dao

type User struct {
	IdUsuario     int     `gorm:"primary_key;column:Id_usuario;autoIncrement"`
	NombreUsuario string  `gorm:"column:Nombre_Usuario; not null; unique"`
	Nombre        string  `gorm:"column:Nombre; not null"`
	Apellido      string  `gorm:"column:Apellido; not null"`
	Email         string  `gorm:"column:Email; not null"`
	Contrasena    string  `gorm:"column:Contraseña; not null"`
	Tipo          bool    `gorm:"column:Tipo; not null"`
	Cursos        []Curso `gorm:"many2many:inscripciones;foreignKey:IdUsuario;joinForeignKey:IdUsuario;References:IdCurso;joinReferences:IdCurso"`
}
