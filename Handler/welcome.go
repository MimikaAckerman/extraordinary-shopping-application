package Handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Welcome(c *gin.Context) {
	c.String(http.StatusOK, "¡Bienvenido a la API de la aplicacion de PSC!")
}
