package Handler

import (
	"bk-compras-extraordinarias/config"
	"bk-compras-extraordinarias/models"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginUser(c *gin.Context) {
	//declaramos las variables usando el paquete de models
	var user models.User
	var userDetails models.UserDetails
	var storedPassword string

	// Parsear el cuerpo de la solicitud JSON
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Solicitud inválida"})
		return
	}

	// Consulta SQL
	query := `SELECT PkUserweb, username, password, Nombre, Apellidos 
	          FROM UbUsersweb WHERE username = @username`

	// Ejecutar la consulta
	err := config.DBSecondary.QueryRow(query, sql.Named("username", user.Username)).
		Scan(&userDetails.PkUserweb, &userDetails.Username, &storedPassword, &userDetails.Nombre, &userDetails.Apellidos)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error en el servidor"})
		}
		return
	}

	// //comparar contraseñas usando bcrypt
	// if err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(user.Password)); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Usuario o contraseña incorrectos"})
	// 	return
	// }

	// Login exitoso
	c.JSON(http.StatusOK, gin.H{
		"message":  "Login exitoso",
		"userData": userDetails,
	})
}
