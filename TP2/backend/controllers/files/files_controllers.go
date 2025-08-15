package files

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var uploadDir = "./uploads" // Definir una variable para la ruta del directorio

func UploadFile(c *gin.Context) {
	// Crear el directorio de subida si no existe
	err := os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		// Si hay un error al crear el directorio, devolver una respuesta JSON con el error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating upload directory"})
		return
	}

	// Obtener el archivo del formulario de la solicitud
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		// Si hay un error al obtener el archivo, devolver una respuesta JSON con el error
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving the file"})
		return
	}
	defer file.Close() // Asegurarse de cerrar el archivo al finalizar

	filename := filepath.Base(header.Filename) // Obtener el nombre base del archivo

	// Crear el archivo en el directorio de subida
	out, err := os.Create(filepath.Join(uploadDir, filename))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating the file"})
		return
	}
	defer out.Close() // Asegurarse de cerrar el archivo al finalizar

	// Copiar el contenido del archivo subido al archivo creado
	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error saving the file"})
		return
	}

	// Devolver una respuesta JSON indicando que el archivo se ha subido con éxito
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Successfully Uploaded File: %s", filename)})
}

func ListFiles(c *gin.Context) {
	files := []string{} // Crear un slice para almacenar los nombres de los archivos

	// Recorrer el directorio de subida
	err := filepath.Walk(uploadDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err // Si hay un error, devolverlo
		}
		if !info.IsDir() {
			// Si no es un directorio, agregar el nombre del archivo al slice
			files = append(files, info.Name())
		}
		return nil
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error listing files: %v", err)})
		return
	}

	// Devolver una respuesta JSON con la lista de archivos
	c.JSON(http.StatusOK, files)
}
