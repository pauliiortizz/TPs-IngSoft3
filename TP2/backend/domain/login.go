package domain

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	IdUser int    `json:"id_usuario"` //este me devuelve
	Token  string `json:"token"`
	Tipo   bool   `json:"tipo"`
}

type User struct {
	IdUsuario     int    `json:"id_usuario"`
	NombreUsuario string `json:"nombre_usuario"`
	Nombre        string `json:"nombre"`
	Apellido      string `json:"apellido"`
	Email         string `json:"email"`
	Contrasena    string `json:"contrasena"`
	Tipo          bool   `json:"tipo"`
}
