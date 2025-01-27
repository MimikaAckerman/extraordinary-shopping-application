package Handler

import (
	"bk-compras-extraordinarias/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

// Endpoint para enviar correos electrónicos
func SendEmail(c *gin.Context) {
	var emailReq models.EmailRequest

	// Validar los datos recibidos
	if err := c.ShouldBindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos: " + err.Error()})
		return
	}

	// Enviar el correo
	if err := enviarCorreo(emailReq.To, emailReq.Subject, emailReq.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el correo: " + err.Error()})
		return
	}

	// Respuesta de éxito
	c.JSON(http.StatusOK, gin.H{"message": "Correo enviado correctamente"})
}

// Función para enviar correos
func enviarCorreo(to, subject, body string) error {
	m := gomail.NewMessage()

	// Configuración del correo
	m.SetHeader("From", "cau@grupoub.com") // Remitente
	m.SetHeader("To", to)                  // Destinatario
	m.SetHeader("Subject", subject)        // Asunto
	m.SetBody("text/html", body)           // Cuerpo del correo

	// Configuración del servidor SMTP
	d := gomail.NewDialer("smtp.ionos.es", 587, "cau@grupoub.com", "23FG@*saj$")

	// Enviar el correo
	return d.DialAndSend(m)
}
