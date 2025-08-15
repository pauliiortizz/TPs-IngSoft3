package db

import (
	"backend/dao"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var sqlDB *sql.DB

func InitDB() {

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	//dsn := "root:ladrillo753@tcp(127.0.0.1:3306)/pbbv?charset=utf8mb3&parseTime=True&loc=Local"
	//dsn := "root:belusql1@tcp(127.0.0.1:3306)/pbbv?charset=utf8mb3&parseTime=True&loc=Local"
	//dsn := "root:BMKvr042@tcp(127.0.0.1:3306)/pbbv?charset=utf8mb3&parseTime=True&loc=Local"
	//dsn := "root:RaTa8855@tcp(127.0.0.1:3306)/pbbv?charset=utf8mb3&parseTime=True&loc=Local"
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	// Obtener la conexión SQL nativa de gorm
	sqlDB, err = DB.DB()
	if err != nil {
		log.Fatal("failed to get sql.DB from gorm: ", err)
	}

	// Migrar las estructuras a la base de datos
	Migrate()

}

func Migrate() {

	// Verificar si la tabla User existe
	userTableExists := DB.Migrator().HasTable(&dao.User{})

	// Verificar si la tabla Curso existe
	cursoTableExists := DB.Migrator().HasTable(&dao.Curso{})

	// Verificar si la tabla Inscripciones existe
	inscripcionesTableExists := DB.Migrator().HasTable(&dao.Inscripciones{})

	if !userTableExists && !cursoTableExists && !inscripcionesTableExists {
		// Migrar tablas en el orden correcto

		err := DB.Migrator().CreateTable(&dao.User{})
		if err != nil {
			log.Fatal("failed to migrate User table: ", err)
		}

		err = DB.Migrator().CreateTable(&dao.Curso{})
		if err != nil {
			log.Fatal("failed to migrate Curso table: ", err)
		}

		err = DB.Migrator().CreateTable(&dao.Inscripciones{})
		if err != nil {
			log.Fatal("failed to migrate Inscripcion table: ", err)
		}

		// Seed the database with initial data
		SeedDB()
	}
}

func SeedDB() {
	// Crear usuarios
	users := []dao.User{
		{NombreUsuario: "pauliiortizz", Nombre: "paulina", Apellido: "ortiz", Email: "pauliortiz@example.com", Contrasena: "contraseña1", Tipo: false},
		{NombreUsuario: "baujuncos", Nombre: "bautista", Apellido: "juncos", Email: "baujuncos@example.com", Contrasena: "contraseña2", Tipo: true},
		{NombreUsuario: "belutreachi", Nombre: "belen", Apellido: "treachi", Email: "belutreachi2@example.com", Contrasena: "contraseña3", Tipo: false},
		{NombreUsuario: "virchu", Nombre: "virginia", Apellido: "rodriguez", Email: "virchurodiguez@example.com", Contrasena: "contraseña4", Tipo: false},
		{NombreUsuario: "johndoe", Nombre: "John", Apellido: "Doe", Email: "johndoe@example.com", Contrasena: "contraseña5", Tipo: false},
		{NombreUsuario: "alicesmith", Nombre: "Alice", Apellido: "Smith", Email: "alicesmith@example.com", Contrasena: "contraseña6", Tipo: true},
		{NombreUsuario: "bobjohnson", Nombre: "Bob", Apellido: "Johnson", Email: "bobjohnson@example.com", Contrasena: "contraseña7", Tipo: false},
		{NombreUsuario: "janedoe", Nombre: "Jane", Apellido: "Doe", Email: "janedoe@example.com", Contrasena: "contraseña8", Tipo: false},
		{NombreUsuario: "emilywilliams", Nombre: "Emily", Apellido: "Williams", Email: "emilywilliams@example.com", Contrasena: "contraseña9", Tipo: true},
	}

	for i, user := range users {
		// Hashear la contraseña con SHA-1
		hasher := sha1.New()
		hasher.Write([]byte(user.Contrasena))
		hashedPassword := hex.EncodeToString(hasher.Sum(nil))
		users[i].Contrasena = hashedPassword
		DB.FirstOrCreate(&users[i], dao.User{Email: user.Email})
	}

	// Crear cursos
	cursos := []dao.Curso{
		{Titulo: "Go Pro: Desarrollo Eficiente con Golang", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "cursodeGo.pdf", Descripcion: "El curso de programación en Go está diseñado para enseñar los fundamentos del lenguaje de programación Go (Golang) de Google. Los estudiantes aprenderán la sintaxis de Go, tipos de datos, estructuras de control y funciones. El curso también cubre la programación concurrente con goroutines y canales, manejo de errores, y el uso de paquetes estándar y de terceros. Mediante ejemplos prácticos y proyectos, los participantes podrán desarrollar aplicaciones eficientes y escalables, preparándose para trabajos en desarrollo backend, servicios web y sistemas distribuidos."},
		{Titulo: "Python Total: De Principiante a Experto en Programación", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "curso_python.pdf", Descripcion: "El curso de programación en Python es una introducción completa al lenguaje de programación Python, diseñado para principiantes y programadores con experiencia previa. A lo largo del curso, los estudiantes aprenderán los fundamentos de Python, incluyendo sintaxis básica, estructuras de datos, funciones y módulos. Además, se explorarán conceptos más avanzados como programación orientada a objetos, manejo de excepciones y manipulación de archivos. El curso también incluye proyectos prácticos y ejercicios que permiten a los estudiantes aplicar lo aprendido en situaciones del mundo real. Al finalizar, los participantes estarán capacitados para desarrollar aplicaciones, automatizar tareas y analizar datos utilizando Python."},
		{Titulo: "Dominando Java: De Principiante a Experto", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "curso_java.pdf", Descripcion: "El curso de programación en Java es una introducción detallada al lenguaje de programación Java, adecuado tanto para principiantes como para programadores con experiencia. Los estudiantes aprenderán los conceptos básicos de Java, como la sintaxis, tipos de datos, control de flujo y estructuras de datos. Además, se abordarán temas avanzados como la programación orientada a objetos, la gestión de excepciones, la concurrencia y el acceso a bases de datos mediante JDBC. A través de ejercicios prácticos y proyectos, los participantes desarrollarán aplicaciones robustas y eficientes, preparándose para roles en desarrollo de software y aplicaciones empresariales."},
		{Titulo: "C++: Programación de Alto Rendimiento", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "cpp.pdf", Descripcion: "Este curso de programación en C++ es una guía completa para aprender uno de los lenguajes de programación más poderosos y versátiles. Los estudiantes comenzarán con los fundamentos de C++, incluyendo sintaxis básica, operadores, estructuras de control y funciones. Luego, se adentrarán en la programación orientada a objetos, manejo de memoria, punteros, y el uso de bibliotecas estándar. El curso también incluye la programación avanzada con plantillas, manejo de archivos y técnicas de optimización. Con proyectos prácticos, los participantes estarán equipados para desarrollar software de alto rendimiento y aplicaciones de sistemas."},
		{Titulo: "JavaScript: Desarrollo Web Interactivo", FechaInicio: time.Now(), Categoria: "Programación", Archivo: "curso_js.pdf", Descripcion: "El curso de programación en JavaScript proporciona una base sólida en el lenguaje de programación JavaScript, esencial para el desarrollo web. Los estudiantes aprenderán la sintaxis de JavaScript, manipulación del DOM, eventos y funciones. El curso también cubre conceptos avanzados como promesas, asincronía con async/await, y el uso de APIs del navegador. Además, se exploran frameworks y bibliotecas populares como React y Node.js. A través de proyectos prácticos, los participantes desarrollarán aplicaciones web interactivas y dinámicas, preparándose para roles en desarrollo frontend y full-stack."},
		{Titulo: "Full-Stack Mastery: Desarrollo Web Completo", FechaInicio: time.Now(), Categoria: "Desarrollo Web", Archivo: "Backend.pdf", Descripcion: "Este curso de desarrollo web full-stack cubre todo el espectro del desarrollo web, desde el frontend hasta el backend. Los estudiantes aprenderán a utilizar HTML, CSS y JavaScript para crear interfaces de usuario atractivas. Luego, se abordarán tecnologías backend como Node.js, Express, y bases de datos relacionales y no relacionales (SQL y MongoDB). El curso incluye el uso de frameworks como React o Angular para el frontend, y la implementación de APIs RESTful. Con proyectos integrales, los participantes estarán capacitados para desarrollar aplicaciones web completas y escalables."},
		{Titulo: "Mobile App Mastery: Desarrollo en iOS y Android", FechaInicio: time.Now(), Categoria: "Aplicaciones Móviles", Archivo: "queesunaaplicacionmovil.ppt", Descripcion: "El curso de desarrollo de aplicaciones móviles enseña a crear aplicaciones para dispositivos iOS y Android. Los estudiantes aprenderán los fundamentos de Swift y Kotlin, los lenguajes de programación utilizados para desarrollar en estas plataformas. El curso cubre el diseño de interfaces de usuario, gestión de estados, acceso a APIs y bases de datos, y publicación en tiendas de aplicaciones. A través de proyectos prácticos, los participantes adquirirán las habilidades necesarias para desarrollar y lanzar aplicaciones móviles."},
		{Titulo: "Data Science & Machine Learning: De Cero a Experto", FechaInicio: time.Now(), Categoria: "Machine Learning", Descripcion: "Este curso de ciencia de datos y machine learning proporciona una introducción completa a la manipulación de datos y el aprendizaje automático. Los estudiantes aprenderán a usar Python y bibliotecas como pandas, NumPy, scikit-learn y TensorFlow. El curso cubre técnicas de preprocesamiento de datos, algoritmos de machine learning supervisado y no supervisado, y redes neuronales. Con proyectos prácticos, los participantes desarrollarán modelos predictivos y análisis de datos, preparándose para roles en análisis de datos y desarrollo de IA."},
		{Titulo: "Cybersecurity Essentials: Protege y Defiende", FechaInicio: time.Now(), Categoria: "Ciberseguridad", Descripcion: "El curso de ciberseguridad enseña los principios fundamentales para proteger sistemas informáticos y redes. Los estudiantes aprenderán sobre criptografía, análisis de vulnerabilidades, control de acceso, y prácticas de seguridad en redes y aplicaciones. El curso también cubre la respuesta a incidentes, pruebas de penetración y el cumplimiento de normativas de seguridad. A través de ejercicios prácticos y estudios de caso, los participantes estarán preparados para roles en seguridad informática y defensa contra ciberataques."},
	}
	for _, curso := range cursos {
		DB.FirstOrCreate(&curso, dao.Curso{Titulo: curso.Titulo})
	}

	// Crear inscripciones
	inscripciones := []dao.Inscripciones{
		{IdUsuario: 1, IdCurso: 1, FechaInscripcion: time.Now(), Comentario: "Espero sobrevivir a este curso y no abandonar"},
		{IdUsuario: 2, IdCurso: 2, FechaInscripcion: time.Now(), Comentario: "Soy el profesor, sean respetuosos en sus comentarios gente"},
		{IdUsuario: 3, IdCurso: 3, FechaInscripcion: time.Now(), Comentario: "Aguante TAYLOR SWIFTTT"},
		{IdUsuario: 4, IdCurso: 4, FechaInscripcion: time.Now(), Comentario: ""},
		{IdUsuario: 5, IdCurso: 5, FechaInscripcion: time.Now(), Comentario: ""},
		{IdUsuario: 1, IdCurso: 2, FechaInscripcion: time.Now(), Comentario: "Cuando empezaba el curso?"},
	}
	for _, inscripcion := range inscripciones {
		DB.FirstOrCreate(&inscripcion, dao.Inscripciones{IdUsuario: inscripcion.IdUsuario, IdCurso: inscripcion.IdCurso})
	}
}

func DeleteCursoByID(IdCurso string) error {
	result := DB.Delete(&dao.Curso{}, "Id_curso = ?", IdCurso)
	return result.Error
}

func DeleteInscripcionesByCursoID(cursoID string) error {
	query := `DELETE FROM inscripciones WHERE Id_curso = ?`
	_, err := sqlDB.Exec(query, cursoID)
	if err != nil {
		log.Printf("Error al eliminar inscripciones: %v\n", err)
		return fmt.Errorf("error al eliminar inscripciones: %w", err)
	}
	return nil
}

func GetCommentsByCourseID(courseID int) ([]dao.Inscripciones, error) {
	var inscripciones []dao.Inscripciones
	result := DB.Where("id_curso = ?", courseID).Find(&inscripciones)
	if result.Error != nil {
		log.Printf("Error al obtener comentarios: %v\n", result.Error)
		return nil, result.Error
	}
	return inscripciones, nil
}

func GetUserIDByUsername(username string) (int, error) {
	var user dao.User

	result := DB.Where("Nombre_Usuario = ?", username).First(&user)

	if result.Error != nil {
		return user.IdUsuario, fmt.Errorf("No se encontro el usuario: %s", username)
	}

	/*if err := DB.Where("username = ?", username).First(&user).Error; err != nil {
		return 0, err
	}*/

	return user.IdUsuario, nil
}

func GetCursosUsuario(userId int) ([]dao.Curso, error) {
	// Obtener los IDs de los cursos del usuario desde la tabla de inscripciones
	var inscripciones []struct {
		IdCurso int `gorm:"column:Id_curso"`
	}
	if err := DB.Table("inscripciones").
		Where("Id_usuario = ?", userId).
		Pluck("Id_curso", &inscripciones).
		Error; err != nil {
		log.Printf("Error al obtener inscripciones del usuario: %v\n", err)
		return nil, err
	}

	// Extraer los IDs de los cursos de la lista de inscripciones
	var cursoIDs []int
	for _, inscripcion := range inscripciones {
		log.Printf("Inscripción encontrada: %+v\n", inscripcion)
		cursoIDs = append(cursoIDs, inscripcion.IdCurso)
	}

	if len(cursoIDs) == 0 {
		log.Println("El usuario no está inscrito en ningún curso")
		return []dao.Curso{}, nil
	}

	// Buscar los cursos correspondientes a los IDs obtenidos
	var cursos []dao.Curso
	if err := DB.Where("Id_curso IN ?", cursoIDs).Find(&cursos).Error; err != nil {
		log.Printf("Error al obtener los cursos: %v\n", err)
		return nil, err
	}

	// Verificar que los datos se han obtenido correctamente
	for _, curso := range cursos {
		log.Printf("Curso obtenido: %+v\n", curso)
	}

	return cursos, nil
}

func GetUsuarioByUsername(username string) (dao.User, error) {
	var usuario dao.User

	result := DB.Where("Nombre_Usuario = ?", username).First(&usuario)

	if result.Error != nil {
		return dao.User{}, fmt.Errorf("No se encontro el usuario: %s", username)
	}

	return usuario, nil
}

func FindCoursesByQuery(query string) ([]dao.Curso, error) {
	var cursos []dao.Curso
	result := DB.Where("Titulo LIKE ? OR Categoria LIKE ?", "%"+query+"%", "%"+query+"%").Find(&cursos)
	if result.Error != nil {
		return nil, result.Error
	}
	return cursos, nil
}

func FindCourseByID(id int) (dao.Curso, error) {
	var curso dao.Curso
	result := DB.First(&curso, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return dao.Curso{}, fmt.Errorf("no se encontró el curso con el ID: %d", id)
		}
		return dao.Curso{}, fmt.Errorf("error al buscar el curso: %w", result.Error)
	}
	return curso, nil
} //Busca el curso

func SubscribeUserToCourse(id_usuario int, id_curso int, fecha_inscripcion time.Time, comentario string) error {
	inscripcion := dao.Inscripciones{
		IdUsuario:        id_usuario,
		IdCurso:          id_curso,
		FechaInscripcion: fecha_inscripcion,
		Comentario:       comentario,
	}
	if err := DB.Create(&inscripcion).Error; err != nil {
		log.Printf("Error al guardar la inscripción: %v\n", err)
		return fmt.Errorf("error al guardar la inscripción: %w", err)
	}
	return nil
} //Crea la inscripcion

func SelectUserByID(id int) (dao.User, error) {
	var user dao.User
	result := DB.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return dao.User{}, fmt.Errorf("No se encontro el usuario con el id: %d", id)
		}
		return dao.User{}, fmt.Errorf("error al buscar el usuario: %w", result.Error)
	}
	return user, nil
} //Retorna el user

func GetAllCursos() ([]dao.Curso, error) {
	var cursos []dao.Curso
	if err := DB.Find(&cursos).Error; err != nil {
		return nil, fmt.Errorf("error obteniendo cursos de la base de datos: %w", err)
	}
	return cursos, nil
}

func GetUserTypeByID(userID int) (bool, error) {
	var user dao.User
	result := DB.First(&user, userID)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return false, fmt.Errorf("no se encontró el usuario con el ID: %d", userID)
		}
		return false, fmt.Errorf("error al buscar el usuario: %w", result.Error)
	}
	return user.Tipo, nil
}
