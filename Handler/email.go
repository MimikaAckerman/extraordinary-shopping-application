package Handler

import (
	"fmt"

	"gopkg.in/gomail.v2"
)

// Función genérica para enviar correos
func EnviarCorreo(to, subject, body string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", "cau@grupoub.com") // Remitente
	m.SetHeader("To", to)                  // Destinatario
	m.SetHeader("Subject", subject)        // Asunto
	m.SetBody("text/html", body)           // Cuerpo del correo

	d := gomail.NewDialer("smtp.ionos.es", 587, "cau@grupoub.com", "23FG@*saj$")
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}

// Función para enviar correo al solicitante
func EnviarCorreoSolicitante(titulo, descripcion, proyecto, coste, usuarioCorreo string) error {
	body := fmt.Sprintf(`
		<div>
			<h2>Confirmación de Solicitud</h2>
			<p><strong>Título:</strong> %s</p>
			<p><strong>Descripción:</strong> %s</p>
			<p><strong>Proyecto:</strong> %s</p>
			<p><strong>Coste:</strong> %s</p>
		</div>
	`, titulo, descripcion, proyecto, coste)

	return EnviarCorreo(usuarioCorreo, "Confirmación de Solicitud Registrada", body)
}

// Función para enviar correo al aprobador
func EnviarCorreoAprobador(titulo, descripcion, proyecto, coste, aprobadorCorreo string) error {
	body := fmt.Sprintf(`
		<div>
			<h2>Nueva Solicitud para Aprobar</h2>
			<p><strong>Título:</strong> %s</p>
			<p><strong>Descripción:</strong> %s</p>
			<p><strong>Proyecto:</strong> %s</p>
			<p><strong>Coste:</strong> %s</p>
			<a href="http://192.168.1.104/UB-DATAWAREHOUSE/login.html" style="padding: 10px 20px; background-color: #007bff; color: white; text-decoration: none; border-radius: 5px; font-weight: bold;">
				Revisar Solicitud
			</a>
		</div>
	`, titulo, descripcion, proyecto, coste)

	return EnviarCorreo(aprobadorCorreo, "Nueva Solicitud para Aprobar", body)
}
