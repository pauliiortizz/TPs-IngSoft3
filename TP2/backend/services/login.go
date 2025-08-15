package services

import (
	"backend/db"
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strings"
	"time"
)

var jwtKey = []byte("my_secret_key")

// Claims estructura para los claims del token JWT
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// generateJWT genera un token JWT para un usuario dado
func generateJWT(username string) (string, error) {
	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(72 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

// Login maneja la autenticación de un usuario
func Login(username string, password string) (string, error) {
	if strings.TrimSpace(username) == "" {
		return "", errors.New("username is required")
	}

	if strings.TrimSpace(password) == "" {
		return "", errors.New("password is required")
	}
	//Verifica que todos los campos esten completos

	hash := fmt.Sprintf("%x", sha1.Sum([]byte(password)))
	//Hashea la contraseña

	userDAO, err := db.GetUsuarioByUsername(username)
	if err != nil {
		return "", fmt.Errorf("error getting user from DB: %w", err)
	}
	//Se devuleve error si es que no se encontro el usuario

	if hash != userDAO.Contrasena {
		return "", errors.New("invalid credentials")
	}
	//Comparacion del hash guardado en la BD y el generado

	token, err := generateJWT(username)
	if err != nil {
		return "", fmt.Errorf("error generating JWT token: %w", err)
	}

	return token, nil
}
