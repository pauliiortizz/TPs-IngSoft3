package users

import (
	"backend/db"
	"backend/domain"
	"backend/services"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Login(c *gin.Context) {
	var creds domain.Credentials

	if err := c.BindJSON(&creds); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{Message: fmt.Sprintf("Invalid request: %s", err.Error())})
		return
	}

	token, err := services.Login(creds.Username, creds.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{
			Message: fmt.Sprintf("Unauthorized login: %s", err.Error()),
		})
		return
	}

	id, err := db.GetUserIDByUsername(creds.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{Message: fmt.Sprintf("Failed to get user ID: %s", err.Error())})
		return
	} //Services

	tipo, err := db.GetUserTypeByID(id)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{Message: fmt.Sprintf("Failed to get user type: %s", err.Error())})
		return
	}

	c.JSON(http.StatusOK, domain.LoginResponse{
		IdUser: id,
		Token:  token,
		Tipo:   tipo,
	}) //Services
}
