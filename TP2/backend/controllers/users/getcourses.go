package users

import (
	"backend/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func ListarCursosUsuario(c *gin.Context) {
	// Obtener el ID del usuario de la URL
	userIdStr := c.Param("id_usuario")
	userId, err := strconv.Atoi(userIdStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de usuario no válido"})
		return
	}

	// Consultar la base de datos para obtener los cursos del usuario
	cursosUsuario, err := db.GetCursosUsuario(userId)
	if err != nil {
		log.Printf("Error al obtener los cursos del usuario: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener los cursos del usuario"})
		return
	}

	// Devolver los cursos encontrados como respuesta
	c.JSON(http.StatusOK, cursosUsuario)
}
