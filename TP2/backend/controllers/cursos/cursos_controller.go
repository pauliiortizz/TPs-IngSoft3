package cursos

import (
	cursosDomain "backend/dao"
	"backend/domain"
	"backend/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func GetComments(c *gin.Context) {
	// Obtiene el parámetro de ruta 'id' del curso
	idCursoStr := c.Param("id")

	// Convierte el parámetro de ruta 'id' de string a int
	idCurso, err := strconv.Atoi(idCursoStr)
	if err != nil {
		// Si hay un error en la conversión, devuelve un error de solicitud incorrecta
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID de curso inválido"})
		return
	}

	// Llama al servicio para obtener los comentarios del curso con el ID especificado
	comentarios, err := services.GetCommentsByCourseID(idCurso)
	if err != nil {
		// Si hay un error al obtener los comentarios, devuelve un error interno del servidor
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Crea un slice para almacenar solo los comentarios no nulos
	var soloComentarios []string
	// Itera sobre los comentarios obtenidos
	for _, comentario := range comentarios {
		// Si el comentario no es una cadena vacía (nulo), lo agrega al slice
		if comentario.Comentario != "" {
			soloComentarios = append(soloComentarios, comentario.Comentario)
		}
	}

	// Devuelve una respuesta JSON con los comentarios no nulos
	c.JSON(http.StatusOK, gin.H{"comentarios": soloComentarios})
}

func DeleteCurso(c *gin.Context) {
	cursoID := c.Param("id")

	err := services.DeleteCurso(cursoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Curso eliminado correctamente"})
}
func UpdateCurso(c *gin.Context) {
	cursoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID del curso inválido"})
		return
	}

	var updatedCurso cursosDomain.Curso
	if err := c.ShouldBindJSON(&updatedCurso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos de curso inválidos"})
		return
	}

	if err := services.UpdateCurso(cursoID, updatedCurso); err != nil {
		log.Printf("Error al actualizar el curso: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar el curso"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Curso editado correctamente"})
}

func CreateCurso(c *gin.Context) {
	var curso cursosDomain.Curso
	if err := c.ShouldBindJSON(&curso); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.CreateCurso(curso); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Curso creado correctamente"})
}

func Search(c *gin.Context) {
	query := strings.TrimSpace(c.Query("query"))
	results, err := services.Search(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Message: fmt.Sprintf("Error in search: %s", err.Error()),
		})
		return
	}

	if len(results) == 0 {
		c.JSON(http.StatusNotFound, domain.Response{
			Message: "No results found",
		})
		return
	}

	c.JSON(http.StatusOK, domain.SearchResponse{
		Results: results,
	})
}

func Get(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Message: fmt.Sprintf("Invalid ID: %s", err.Error()),
		})
		return
	}

	course, err := services.Get(int(id))
	if err != nil {
		c.JSON(http.StatusNotFound, domain.Response{
			Message: fmt.Sprintf("Error in get: %s", err.Error()),
		})
		return
	}

	c.JSON(http.StatusOK, course)
}

func GetAllCursos(c *gin.Context) {
	cursos, err := services.GetAllCursos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, cursos)
}
